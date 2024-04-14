package global

import "daydayup/db/dao"

type GetSummaryReq struct {
	ID string `json:"user_id" binding:"required"`
}

type SignSummaryVO struct {
	ID     string `json:"id" `
	Sign   bool   `json:"sign_today" `
	Streak int    `json:"longest_streak" `
}

func NewSignSummaryVO(record dao.DailySignModel) SignSummaryVO {
	return SignSummaryVO{
		ID:     record.UserID,
		Sign:   record.SignToday == 1,
		Streak: record.Streak,
	}
}

type SignReq struct {
	ID string `json:"user_id" binding:"required"`
}

// todo 暂时不开放
type SignDetailVO struct {
	ID     string `json:"id"`
	Sign   bool   `json:"sign_today" `
	Streak int    `json:"longest_streak" `
}
