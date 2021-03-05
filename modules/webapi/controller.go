package webapi

import (
	"errors"
	"golang-started/dto"
	"golang-started/httperror"
	"golang-started/lib/opentracing"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// Controller 控制器
type Controller struct {
	Service ServiceInterface
}

// CreateDto Post 提交的也可以使用绑定信息做校验
type CreateDto struct {
	// Name 名称
	Name string `json:"name" binding:"required"`
}

// UpdateDto UpdateDto
type UpdateDto struct {
	// 名称
	Name string `json:"name" binding:"required"`
}

// Create 创建
// @Summary 创建
// @Description 创建
// @Param example body example.CreateDto true "创建"
// @Param uid header string true "用户ID"
// @Accept json
// @Router /example/ [post]
// @Success 200 {object} example.Example
func (that *Controller) Create(ctx *gin.Context) {
	h := dto.Header{}
	if err := ctx.ShouldBindHeader(&h); err != nil {
		panic(httperror.BadRequest(err.Error(), 1000))
	}

	var dto CreateDto
	if err := ctx.Bind(&dto); err != nil {
		panic(httperror.BadRequest(err.Error(), 1000))
	}
	m := &wpPosts{}
	copier.Copy(&dto, m)
	_, m, err := that.Service.Create(ctx, m)
	if err != nil {
		panic(httperror.InternalError(err.Error(), 1000))
	}
	ctx.JSON(http.StatusOK, m)
}

func (that *Controller) TestLinkTrace(ctx *gin.Context) {
	//c := &http.Client{}
	ctxNew := ctx.Request.Context()
	httpReq, _ := http.NewRequest("GET", "http://localhost:3002/health", nil)
	opentracing.DoWithLinkTrace(ctxNew, httpReq)
	ctx.JSON(http.StatusOK, "ok")
}

// GetDto 获取课件详情参数
type GetDto struct {
	ID uint64 `uri:"id" binding:"required"`
}

// GetOne 获取单条
// @Summary 获取单条
// @Router /example/{id} [get]
// @Param id path string true "ID"
// @Param uid header string true "uID"
// @Success 200 {object} example.Example
func (that *Controller) GetOne(ctx *gin.Context) {
	dto := &GetDto{}
	err := ctx.BindUri(dto)
	if err != nil {
		panic(err)
	}
	// 调用service
	_, m, err := that.Service.Get(ctx, dto.ID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		panic(httperror.BadRequest("文章内容不存在", 1001))
	}
	if err != nil {
		panic(httperror.InternalError(err.Error(), 1002))
	}
	ctx.JSON(http.StatusOK, m)
}

// 获取分类目录
func (that *Controller) GetCategory(ctx *gin.Context) {
	// 调用service
	_, m, err := that.Service.GetCategory(ctx)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		panic(httperror.BadRequest("分类目录不存在", 1001))
	}
	if err != nil {
		panic(httperror.InternalError(err.Error(), 1002))
	}
	data := struct {
		Code int        `json:"code"`
		Msg  string     `json:"msg"`
		List []category `json:"list"`
	}{
		Code: 0,
		Msg:  "success",
		List: m,
	}
	ctx.JSON(http.StatusOK, data)
}
