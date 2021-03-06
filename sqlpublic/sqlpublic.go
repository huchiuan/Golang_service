package sqlbublic

import (
	"fmt"
	"time"

	"os"

	Users "UI_Assignment/model"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

func SqlInit() {
	e := godotenv.Load()
	if e != nil {
		log.Error(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	createClient(username, password, dbName, dbHost)
}

func createClient(username string, password string, dbName string, dbHost string) {

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s  password=%s ", dbHost, username, dbName, password)

	log.Debug(dbUri)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Error("<SqlPublic>:SQL連接失敗\n")
		log.Panic(err)

	}

	db = conn

	sqlDB, err := db.DB()
	if err != nil {
		log.Panic(err)

	}

	sqlDBStats := sqlDB.Stats()
	errword := sqlDB.Ping()

	log.Debug("sqlDBStats")
	log.Debug(sqlDBStats)
	log.Debug("\n")
	log.Debug(errword)
	log.Debug("\n")

	log.Info("<SqlPublic>:SQL連接成功 \n")

}

func Listallusers() ([]Users.User, error) {
	// var users Users.Users
	var user []Users.User

	result := db.Find(&user)
	log.Debug(result)

	if result.Error == nil {
		log.Info("<SqlPublic>:查詢結果:", result)
		return user, nil
	} else {
		log.Error("<SqlPublic>:查詢失敗")
		return user, nil
	}

}

func Searchanuser(fullname string) (Users.User, error) {
	// var users Users.Users
	var user Users.User

	result := db.First(&user, "fullname = ?", fullname)
	log.Debug(result)

	if result.Error == nil {
		log.Info("<SqlPublic>:查詢結果:", result)
		return user, nil
	} else {
		log.Error("<SqlPublic>:查詢失敗")
		return user, nil
	}

}

func Getuserdetailedinformation() ([]Users.User, error) {
	// var users Users.Users
	var user []Users.User

	result := db.Find(&user)
	log.Debug(result)

	if result.Error == nil {
		log.Info("<SqlPublic>:查詢結果:", result)
		return user, nil
	} else {
		log.Error("<SqlPublic>:查詢失敗")
		return user, nil
	}

}

func Createtheuser(user Users.User) error {

	result := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Table("users").Create(&user)

	if result.Error == nil {
		log.Info("<SqlPublic>:新增資料成功\n")

		return nil
	} else {

		log.Error("<SqlPublic>:新增資料失敗\n")

		return result.Error
	}

}

func Getuserbyacct(acct string) (Users.User, error) {

	var user Users.User

	result := db.First(&user, "acct = ?", acct)
	log.Debug(result)

	if result.Error == nil {
		log.Info("<SqlPublic>:查詢結果:", result)
		return user, nil
	} else {
		log.Error("<SqlPublic>:查詢失敗")
		return user, nil
	}

}

func Deleteuserbyacct(acct string) error {

	var user Users.User
	user.Acct = acct
	result := db.Delete(&user)
	log.Debug(result)

	if result.Error == nil {
		log.Info("<SqlPublic>:刪除成功")
		return nil
	} else {
		log.Error("<SqlPublic>:刪除失敗")
		return nil
	}

}

func Updateuser(acct, pwd string) error {

	var user Users.User
	user.Acct = acct
	result := db.First(&user)
	if result.Error == nil {

		user.Pwd = pwd
		user.Updated_at = time.Now()
		result = db.Save(&user)

		if result.Error == nil {
			log.Info("<SqlPublic>:更新結束")
			return nil
		} else {
			log.Error("<SqlPublic>:更新失敗")
			return nil
		}

	} else {
		log.Error("<SqlPublic>:更新失敗，查無此人")
		return nil
	}

}

func Updateuserfullname(acct, fullname string) error {

	var user Users.User
	user.Acct = acct
	result := db.First(&user)
	if result.Error == nil {

		user.Fullname = fullname
		user.Updated_at = time.Now()
		result = db.Save(&user)

		if result.Error == nil {
			log.Info("<SqlPublic>:更新結束")
			return nil
		} else {
			log.Error("<SqlPublic>:更新失敗")
			return nil
		}

	} else {
		log.Error("<SqlPublic>:更新失敗，查無此人")
		return nil
	}

}
