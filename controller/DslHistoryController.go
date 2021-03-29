package controller

import (
	"ElasticView/model"
	"ElasticView/platform-basic-libs/jwt"
	"ElasticView/platform-basic-libs/response"

	"github.com/gin-gonic/gin"
)

type DslHistoryController struct {
	BaseController
}

func (this DslHistoryController) ListAction(ctx *gin.Context) {
	c, err := jwt.ParseToken(ctx.GetHeader("X-Token"))
	if err != nil {
		this.Error(ctx, err)
		return
	}
	gmDslHistoryModel := model.GmDslHistoryModel{}
	err = ctx.Bind(&gmDslHistoryModel)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	gmDslHistoryModel.Uid = int(c.ID)

	list, err := gmDslHistoryModel.List()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	count, err := gmDslHistoryModel.Count()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": list, "count": count})
}

func (this DslHistoryController) CleanAction(ctx *gin.Context) {
	c, err := jwt.ParseToken(ctx.GetHeader("X-Token"))
	if err != nil {
		this.Error(ctx, err)
		return
	}
	gmDslHistoryModel := model.GmDslHistoryModel{}

	gmDslHistoryModel.Uid = int(c.ID)
	err = gmDslHistoryModel.Clean()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
}
