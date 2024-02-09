package handler

import (
	"net/http"
	"simplink/features/links/core"
	"simplink/helpers/exceptions"

	echo "github.com/labstack/echo/v4"
)

func NewLinkHandler(service core.Service) core.Handler {
	return &linkHandler{
		service: service,
	}
}

type linkHandler struct {
	service core.Service
}

func (hdl *linkHandler) GetByShort(ctx echo.Context) error {
	GetResponse := make(DefaultResponse)

	short := ctx.Param("shorten-link")
	data, err := hdl.service.GetByShort(ctx.Request().Context(), short)
	if err != nil {
		switch err.(type) {
		case exceptions.Validation:
			GetResponse["message"] = err.Error()
			return ctx.JSON(http.StatusBadRequest, GetResponse)
		case exceptions.Repository:
			GetResponse["message"] = err.Error()
			return ctx.JSON(http.StatusNotFound, GetResponse)
		}

		return err
	}

	return ctx.Redirect(302, data.Destination)
}

func (hdl *linkHandler) Create(ctx echo.Context) error {
	CreateResponse := make(DefaultResponse)
	CreateRequest := new(LinkCreateRequest)

	if err := ctx.Bind(CreateRequest); err != nil {
		return err
	}

	result, err := hdl.service.Create(ctx.Request().Context(), CreateRequest.ToEntity())
	if err != nil {
		switch err.(type) {
		case exceptions.Validation:
			CreateResponse["message"] = err.Error()
			return ctx.JSON(http.StatusBadRequest, CreateResponse)
		case exceptions.Repository:
			CreateResponse["message"] = err.Error()
			return ctx.JSON(http.StatusConflict, CreateResponse)
		}

		return err
	}

	CreateResponse["message"] = "link shortened successfully"

	if result != nil {
		CreateResponse["data"] = LinkResponseFromEntity(*result)
	}

	return ctx.JSON(http.StatusCreated, CreateResponse)
}
