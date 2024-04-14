package dao

import (
	"database/sql"
	"daydayup/db"
	"errors"
	"log"
	"time"
)

type DailySignModel struct {
	ID        int    `json:"id,omitempty" db:"id,omitempty"`
	UserID    string `json:"user_id" db:"user_id"`
	January   int32  `json:"january" db:"january"`
	February  int32  `json:"february" db:"february"`
	March     int32  `json:"march" db:"march"`
	April     int32  `json:"april" db:"april"`
	May       int32  `json:"may" db:"may"`
	June      int32  `json:"june" db:"june"`
	July      int32  `json:"july" db:"july"`
	August    int32  `json:"august" db:"august"`
	September int32  `json:"september" db:"september"`
	October   int32  `json:"october" db:"october"`
	November  int32  `json:"november" db:"november"`
	December  int32  `json:"december" db:"december"`
	Streak    int    `json:"streak" db:"streak"`
	SignToday int8   `json:"sign_today" db:"sign_today"`
}

const (
	hasSignedToday = 1
	notSignedYet   = 0

	insertDailySign = `INSERT INTO daily_sign(user_id, january, february, march, april, may, june, july, august, 
                       september, october, november, december, streak, sign_today) VALUES(:user_id,:january,:february,
                       :march,:april,:may,:june,:july,:august,:september,:october,:november,:december,:streak,:sign_today);`
	selectDailySign    = `SELECT * FROM daily_sign WHERE user_id = ?;`
	selectAllDailySign = `SELECT * FROM daily_sign;`
	updateDailySign    = `UPDATE daily_sign SET january=:january, february=:february, march=:march,april=:april,
                      may=:may,june=:june,july=:july,august=:august,september=:september,october=:october,
                      november=:november,december=:december,streak=:streak,sign_today=:sign_today WHERE user_id = :user_id;`
)

type DailySign struct {
	db *db.SQLite
}

func NewDailySignDAO(database db.SQLite) DailySign {
	return DailySign{db: &database}
}
func (ds *DailySign) GetForUID(uid string) (interface{}, error) {
	var res DailySignModel
	if err := ds.db.Get(&res, selectDailySign, uid); err != nil {
		return nil, err
	}
	return res, nil
}

func (ds *DailySign) Insert(record DailySignModel) error {
	_, err := ds.db.NamedExec(insertDailySign, record)
	if err != nil {
		return err
	}
	return nil
}

func (ds *DailySign) Update(model DailySignModel) error {
	_, err := ds.db.NamedExec(updateDailySign, model)
	if err != nil {
		return err
	}
	return nil
}

func (ds *DailySign) ListAll() []DailySignModel {
	var res []DailySignModel
	err := ds.db.Select(&res, selectAllDailySign)
	if err != nil {
		log.Println(err.Error())
	}
	return res
}

func (ds *DailySign) UpdateSignTodayForUID(uid string) (interface{}, error) {
	var res DailySignModel
	err := ds.db.Get(&res, selectDailySign, uid)
	var isNewUser bool
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		isNewUser = true
	} else if err != nil {
		return nil, err
	}
	res.UserID = uid
	if res.SignToday == hasSignedToday {
		return nil, errors.New("already signed today, cannot sign again")
	}
	res.SignToday = hasSignedToday
	res.Streak++

	month := time.Now().Month()
	switch month {
	case time.January:
		res.January = sign(res.January)
	case time.February:
		res.February = sign(res.February)
	case time.March:
		res.March = sign(res.March)
	case time.April:
		res.April = sign(res.April)
	case time.May:
		res.May = sign(res.May)
	case time.June:
		res.June = sign(res.June)
	case time.July:
		res.July = sign(res.July)
	case time.August:
		res.August = sign(res.August)
	case time.September:
		res.September = sign(res.September)
	case time.October:
		res.October = sign(res.October)
	case time.November:
		res.November = sign(res.November)
	case time.December:
		res.December = sign(res.December)
	default:
		return nil, errors.New("undefined month")
	}
	if isNewUser {
		if err := ds.Insert(res); err != nil {
			return nil, err
		}
		return res, nil
	}
	_, err = ds.db.NamedExec(updateDailySign, res)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func sign(record int32) int32 {
	day := time.Now().Day()
	return record | (1 << day)
}

func ResetStreak(db *db.SQLite) {
	ds := NewDailySignDAO(*db)
	records := ds.ListAll()
	for _, record := range records {
		if record.SignToday == notSignedYet {
			record.Streak = 0
		}
		record.SignToday = notSignedYet
		if err := ds.Update(record); err != nil {
			log.Println(err.Error())
		}
	}
	return
}
