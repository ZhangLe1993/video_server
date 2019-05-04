package orm

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ReadVideoDeletionRecord(count int) ([]string, error) {

	stmt, err := dbConn.Prepare("SELECT video_id FROM video_del_rec limit ?")
	var ids []string
	if err != nil {
		log.Printf("Prepare SQL Error %v \n", err)
		return ids, err
	}

	rows, err := stmt.Query(count)
	if err != nil {
		log.Printf("Query video deletion record error, %v \n", err)
		return ids, nil
	}

	for rows.Next() {
		var id string
		if err = rows.Scan(&id); err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	defer stmt.Close()

	return ids, nil

}

func DelVideoDeletionRecord(vid string) error {
	stmt, err := dbConn.Prepare("DELETE FROM video_del_rec WHERE video_id = ?;")
	if err != nil {
		log.Printf("Prepare Delete Video Del Record Error %v \n", err)
		return err
	}
	_, err = stmt.Exec(vid)
	if err != nil {
		log.Printf("Delete Video Del Record Error %v \n", err)
		return err
	}

	defer stmt.Close()
	return nil
}
