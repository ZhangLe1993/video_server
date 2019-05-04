package orm

import "log"

func AddVideoDeletionRecord(vid string) error {
	stmt, err := dbConn.Prepare("INSERT INTO video_rel_rec (video_id) values (?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(vid)
	if err != nil {
		log.Printf("INSERT Video Del Rec Error: %v \n", err)
		return err
	}

	defer stmt.Close()
	return nil
}
