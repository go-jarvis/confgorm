package __example__

import (
	"testing"

	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

func userProcessAll(t *testing.T, db *gorm.DB) {

	autoMigrateUser(db)
	defer dropTableUser(db)

	// select by phone
	cellphone := "19912341234"
	info2 := UserInfo{
		Name:      "zhangsan",
		Age:       18,
		Cellphone: "19912341234",
	}

	t.Run("create user", func(t *testing.T) {
		// create
		user1 := User{
			UserInfo: UserInfo{
				Name:      "zhangsan",
				Age:       18,
				Cellphone: "13312341234",
			},
		}
		err := createUser(db, &user1)
		NewWithT(t).Expect(err).To(BeNil())
		err = createUser(db, &user1)
		NewWithT(t).Expect(err).NotTo(BeNil())

		user2 := User{
			UserInfo: info2,
		}
		err2 := createUser(db, &user2)
		NewWithT(t).Expect(err2).To(BeNil())
	})

	t.Run("select user", func(t *testing.T) {
		_info := getUserByCellphone(db, cellphone)
		NewWithT(t).Expect(_info).To(Equal(info2))
	})

	t.Run("update user", func(t *testing.T) {
		info2 := info2
		info2.Name = "zhugeliang"
		// update by phone
		_info2 := updateUserByCellphone(db, cellphone, "zhugeliang")
		NewWithT(t).Expect(_info2).To(Equal(info2))
	})

	t.Run("delete user", func(t *testing.T) {
		// Delete by phone
		err3 := deleteUserByCellphone(db, cellphone)
		NewWithT(t).Expect(err3).To(BeNil())

	})

}
