package mysql

func (my *MysqlDriver) Magrate() {

	// nerver do action when database target is not same
	if my.Name() != my.Database.Name() {
		return
	}

	for _, table := range my.Database.Tables() {
		my.AutoMigrate(table)
	}
}
