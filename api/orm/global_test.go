package orm

import "testing"

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	//clearTables()
}

func clearTables() {
	conn.Exec("truncate user")
	conn.Exec("truncate video_info")
	conn.Exec("truncate comments")
	conn.Exec("truncate sessions")
}
