package handler

import (
	"fmt"
	"net/http"

	"example.com/test/model"
	"example.com/test/repository"
	"example.com/test/repository/repo_impl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler struct {
	Repo repository.Repo
}

func NewHandler(repo repository.Repo) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) HandlerSum(c echo.Context) error {
	repo := repo_impl.NewRepo()
	h.Repo = repo
	if h.Repo == nil {
		return c.String(http.StatusInternalServerError, "repository is not initialized")
	}
	req := model.Input{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.String(http.StatusBadRequest, "Can't get value")
	}
	res, err := h.Repo.Process(c.Request().Context(), req)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusBadRequest, "Can't process value")
	}
	// return c.String(http.StatusOK, fmt.Sprintf("xin chao ban,nam nay ban: %v", res))

	name := req.Name // Trích xuất tên từ req
	age := res.Age   // Trích xuất tuổi từ res

	return c.String(http.StatusOK, fmt.Sprintf("Xin chào bạn: %s\nNăm nay bạn: %d", name, age))
}
