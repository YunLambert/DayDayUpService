package handler

import (
	"daydayup/db/dao"
	"daydayup/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DailySign struct {
	Dao dao.DailySign
}

func (ds *DailySign) GetSummary(c *gin.Context) {
	var req global.GetSummaryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, global.ErrorWithMsg(http.StatusBadRequest, err.Error()))
		return
	}
	res, err := ds.Dao.GetForUID(req.ID)
	if err != nil {
		c.JSON(http.StatusOK, global.ErrorWithMsg(http.StatusBadRequest, err.Error()))
		return
	}
	c.JSON(http.StatusOK, global.SuccessWithData(res))
	return
}

func (ds *DailySign) Sign(c *gin.Context) {
	var req global.SignReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, global.ErrorWithMsg(http.StatusBadRequest, err.Error()))
		return
	}

	record, err := ds.Dao.UpdateSignTodayForUID(req.ID)
	if err != nil {
		c.JSON(http.StatusOK, global.ErrorWithMsg(http.StatusBadRequest, err.Error()))
		return
	}
	c.JSON(http.StatusOK, global.SuccessWithData(global.NewSignSummaryVO(record.(dao.DailySignModel))))
	return
}
