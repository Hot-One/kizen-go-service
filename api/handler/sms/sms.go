package sms

import (
	statushttp "github.com/Hot-One/kizen-go-service/api/status_http"
	"github.com/Hot-One/kizen-go-service/config"
	"github.com/Hot-One/kizen-go-service/dto"
	"github.com/Hot-One/kizen-go-service/pkg/logger"
	_ "github.com/Hot-One/kizen-go-service/pkg/pg"
	"github.com/Hot-One/kizen-go-service/storage/repo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	cfg     *config.Config
	log     logger.Logger
	storage repo.SmsInterface
}

func NewHandler(group *gin.RouterGroup, cfg *config.Config, log logger.Logger, storage repo.SmsInterface) {
	var h = &handler{
		cfg:     cfg,
		log:     log,
		storage: storage,
	}

	sms := group.Group("/sms")
	{
		sms.POST("/send", h.Create)
		sms.GET("/verify/:id", h.Verify)
	}
}

// Create 		godoc
// @ID			send-sms
// @Tags		Sms
// @Summary		send-sms
// @Description	send-sms
// @Accept		json
// @Produce		json
// @Param		input body dto.CreateSms true "Request body"
// @Success 	201 {object} pg.Id "Successful operation"
// @Failure 	400 {object} statushttp.Response "Bad Request"
// @Failure 	500 {object} statushttp.Response "Internal Server Error"
// @Router		/sms/send [POST]
func (h *handler) Create(c *gin.Context) {
	var input dto.CreateSms

	{
		if err := c.ShouldBindJSON(&input); err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	id, err := h.storage.Create(&input)
	if err != nil {
		statushttp.InternalServerError(c, err.Error())
		return
	}

	statushttp.Created(c, id)
}

// Verify 		godoc
// @ID			verify-sms
// @Tags		Sms
// @Summary		verify-sms
// @Description	verify-sms
// @Accept		json
// @Produce		json
// @Param       id path int64 true "Id"
// @Param       code query string true "Code"
// @Success 	200 {object} statushttp.Response "Successful operation"
// @Failure 	400 {object} statushttp.Response "Bad Request"
// @Failure 	500 {object} statushttp.Response "Internal Server Error"
// @Router		/sms/verify/{id} [GET]
func (h *handler) Verify(c *gin.Context) {
	id, err := statushttp.GetId(c)
	{
		if err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	var queryParams dto.VerifySms
	{
		if err := c.ShouldBindQuery(&queryParams); err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	queryParams.Id = id

	var filter = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ? AND code = ?", queryParams.Id, queryParams.Code)
	}

	_, err = h.storage.FindOne(c.Request.Context(), filter)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			statushttp.BadRequest(c, "invalid code")
			return
		}

		statushttp.InternalServerError(c, err.Error())
		return
	}

	statushttp.OK(c, "code verified")
}
