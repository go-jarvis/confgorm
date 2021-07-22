package __example__

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserInfo `gorm:"embedded"`
}

type UserInfo struct {
	Name      string `json:"name" gorm:"type:varchar(20)"`
	Age       int    `json:"age"`
	Cellphone string `json:"cellphone" gorm:"type:varchar(20);index:;unique;"`
}

func autoMigrateUser(db *gorm.DB) {
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
}
func dropTableUser(db *gorm.DB) {
	if err := db.Migrator().DropTable(&User{}); err != nil {
		panic(err)
	}
}

func createUser(db *gorm.DB, user *User) error {
	ret := db.Create(user)
	if ret.RowsAffected != 1 {
		return ret.Error
	}
	return nil
}

func getUserByCellphone(db *gorm.DB, cellphone string) UserInfo {

	user := User{}

	// db.First(dest interface{}, conds ...interface{})
	db.First(&user, "cellphone = ?", cellphone)

	return user.UserInfo
}

func updateUserByCellphone(db *gorm.DB, cellphone string, name string) UserInfo {

	db.Model(&User{}).Where("cellphone = ?", cellphone).Update("Name", name)

	return getUserByCellphone(db, cellphone)
}

func deleteUserByCellphone(db *gorm.DB, cellphone string) error {
	// db.Delete(&User{}, "cellphone = ?", cellphone)

	ret := db.Where("cellphone = ?", cellphone).Delete(&User{})

	if ret.Error != nil {
		return fmt.Errorf("delete error failed:%v", ret.Error)
	}

	return nil
}
