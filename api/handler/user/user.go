package user

import (
	statushttp "github.com/Hot-One/kizen-go-service/api/status_http"
	"github.com/Hot-One/kizen-go-service/config"
	"github.com/Hot-One/kizen-go-service/dto"
	"github.com/Hot-One/kizen-go-service/pkg/logger"
	"github.com/Hot-One/kizen-go-service/storage/repo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	cfg     config.Config
	log     logger.Logger
	storage repo.UserInterface
}

func NewHandler(group *gin.RouterGroup, cfg config.Config, log logger.Logger, storage repo.UserInterface) {
	var h = &handler{
		cfg:     cfg,
		log:     log,
		storage: storage,
	}

	user := group.Group("/user")
	{
		user.POST("/", h.Create)
		user.GET("/:id", h.Get)
		user.GET("/", h.GetAll)
		user.PUT("/:id", h.Update)
		user.DELETE("/:id", h.Delete)
	}
}

// Create 		godoc
// @ID			create-user
// @Tags		User
// @Summary		create a user
// @Description	create a user
// @Accept		json
// @Produce		json
// @Param		input body dto.CreateUser true "Request body"
// @Success 	201 {object} pg.Id "Successful operation"
// @Failure 	400 {object} statushttp.Response "Bad Request"
// @Failure 	500 {object} statushttp.Response "Internal Server Error"
// @Router		/user [POST]
func (h *handler) Create(c *gin.Context) {
	var input dto.CreateUser

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

// Get 		godoc
// @ID			get-user
// @Tags		User
// @Summary		get a user
// @Description	get a user
// @Accept		json
// @Produce		json
// @Param       id path int64 true "Id"
// @Success 	200 {object} dto.User "Successful operation"
// @Failure 	400 {object} statushttp.Response "Bad Request"
// @Failure 	404 {object} statushttp.Response "Not Found"
// @Failure 	500 {object} statushttp.Response "Internal Server Error"
// @Router		/user/{id} [GET]
func (h *handler) Get(c *gin.Context) {
	id, err := statushttp.GetId(c)
	{
		if err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	var filter = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	user, err := h.storage.FindOne(c.Request.Context(), filter)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			statushttp.NotFound(c, "user not found")
			return
		}
		statushttp.InternalServerError(c, err.Error())
		return
	}

	statushttp.OK(c, user)
}

// GetAll 		godoc
// @ID			get-all-users
// @Tags		User
// @Summary		get all users
// @Description	get all users
// @Accept		json
// @Produce		json
// @Param       page query int false "Page"
// @Param       size query int false "Size"
// @Success 	200 {object} dto.UserPage "Successful operation"
// @Failure 	500 {object} statushttp.Response "Internal Server Error"
// @Router		/user [GET]
func (h *handler) GetAll(c *gin.Context) {
	page, size, err := statushttp.GetPageLimit(c)
	if err != nil {
		statushttp.BadRequest(c, err.Error())
		return
	}

	users, err := h.storage.Page(c.Request.Context(), nil, int64(page), int64(size))
	if err != nil {
		statushttp.InternalServerError(c, err.Error())
		return
	}

	statushttp.OK(c, users)
}

// Update 		godoc
// @ID			update-user
// @Tags		User
// @Summary		update a user
// @Description	update a user
// @Accept		json
// @Produce		json
// @Param       id path int64 true "Id"
// @Param		input body dto.UpdateUser true "Request body"
// @Success 	200 {object} statushttp.Response "Successful operation"
// @Failure 	400 {object} statushttp.Response "Bad Request"
// @Failure 	404 {object} statushttp.Response "Not Found"
// @Failure 	500 {object} statushttp.Response "Internal Server Error"
// @Router		/user/{id} [PUT]
func (h *handler) Update(c *gin.Context) {
	id, err := statushttp.GetId(c)
	{
		if err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	var input dto.UpdateUser
	{
		if err := c.ShouldBindJSON(&input); err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	var filter = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	err = h.storage.Update(&input, filter)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			statushttp.NotFound(c, "user not found")
			return
		}
		statushttp.InternalServerError(c, err.Error())
		return
	}

	statushttp.OK(c, "user updated")
}

// Delete 		godoc
// @ID			delete-user
// @Tags		User
// @Summary		delete a user
// @Description	delete a user
// @Accept		json
// @Produce		json
// @Param       id path int64 true "Id"
// @Success 	200 {object} statushttp.Response "Successful operation"
// @Failure 	400 {object} statushttp.Response "Bad Request"
// @Failure 	404 {object} statushttp.Response "Not Found"
// @Failure 	500 {object} statushttp.Response "Internal Server Error"
// @Router		/user/{id} [DELETE]
func (h *handler) Delete(c *gin.Context) {
	id, err := statushttp.GetId(c)
	{
		if err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	var filter = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	err = h.storage.Delete(filter)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			statushttp.NotFound(c, "user not found")
			return
		}
		statushttp.InternalServerError(c, err.Error())
		return
	}

	statushttp.OK(c, "user deleted")
}
