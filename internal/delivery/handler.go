package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jamascrorpJS/eBank/internal/domain"
	"github.com/jamascrorpJS/eBank/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) NewEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(
		gin.Logger(),
		gin.Recovery(),
		CheckHeader(),
		Digest,
	)
	h.StartServer(engine)
	return engine
}

func (h *Handler) StartServer(engine *gin.Engine) {
	routerGroup := engine.Group("/api")
	{
		h.Routes(routerGroup)
	}
}

func (h *Handler) Routes(r *gin.RouterGroup) {
	{
		r.POST("/operations", h.GetAllOperations)
		r.POST("/exists", h.Existed)
		r.POST("/create", h.CreateUser)

		r.POST("/balance", h.GetBalance)
		r.POST("/update", h.UpdateBalance)
	}
}
func (h *Handler) GetAllOperations(c *gin.Context) {
	body, err := h.service.Operation.GetAllOperations(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(http.StatusOK, body)
}

type Phone struct {
	Phone string `json:"phone" validate:"required,numeric,len=9"`
}

func (h *Handler) Existed(c *gin.Context) {
	var phone Phone
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No valid phone number"})
		return
	}
	validate := validator.New()
	err := validate.Struct(phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}
	body, err := h.service.User.Existed(c.Request.Context(), phone.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(http.StatusOK, body)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user domain.User
	user.ID = uuid.New()
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No valid user"})
		return
	}

	err := h.service.User.CreateUser(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": user})
}

func (h *Handler) GetBalance(c *gin.Context) {
	id := c.MustGet("id")
	userID, ok := id.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	body, err := h.service.Balance.GetBalance(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": body})
}

type Balance struct {
	Sum       float64 `json:"sum" validate:"required,numeric,gte=1"`
	Comission int     `json:"percent" validate:"numeric,gte=0"`
}

func (h *Handler) UpdateBalance(c *gin.Context) {
	id := c.MustGet("id")
	userID, ok := id.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}
	var balance Balance
	if err := c.ShouldBindJSON(&balance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No valid fields"})
		return
	}
	validate := validator.New()
	err := validate.Struct(balance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}
	err = h.service.Balance.UpdateBalance(c.Request.Context(), userID, balance.Sum, balance.Comission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": "updated"})
}
