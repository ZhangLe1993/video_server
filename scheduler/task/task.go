package task

import (
	"errors"
	"log"
	"os"
	"sync"
	"video_server/scheduler/orm"
)

func deleteVideo(vid string) error {
	err := os.Remove("./videos/" + vid)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Delete video error : %v \n", err)
		return err
	}
	return nil
}

func VideoClearDispatch(dc chan interface{}) error {

	res, err := orm.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video Clear Dispatch Error : %v \n", err)
		return err
	}
	if len(res) == 0 {
		return errors.New("All Task Finished")
	}
	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc chan interface{}) error {
	errMap := &sync.Map{}

	var err error

ownLoop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {

				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}

				if err := orm.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)

		default:
			break ownLoop

		}
	}

	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})

	return nil
}
