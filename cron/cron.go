package cron

import (
	"daydayup/db"
	"daydayup/db/dao"
	"github.com/go-co-op/gocron"
	"time"
)

func InitAutoResetStreak(db *db.SQLite) {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(1).Day().At("04:00").Do(dao.ResetStreak, db)
	if err != nil {
		panic(err)
	}
	s.StartAsync()
}
