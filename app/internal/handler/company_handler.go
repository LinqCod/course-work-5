package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/linqcod/course-work-5/app/internal/dto"
	"github.com/linqcod/course-work-5/app/internal/model"
	"net/http"
)

func (h *Handler) CreateCompany(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var company model.Company
	err = c.BindJSON(&company)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateCompany(userId, company)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type AllCompaniesResponse struct {
	Data []model.Company `json:"data"`
}

func (h *Handler) GetAllCompanies(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	companies, err := h.services.Company.GetAllCompanies(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, AllCompaniesResponse{
		Data: companies,
	})
}

func (h *Handler) GetCompanyById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	companyId, err := getId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	company, err := h.services.Company.GetCompanyById(userId, companyId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, company)
}

func (h *Handler) UpdateCompany(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	companyId, err := getId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input dto.UpdateCompanyDto
	err = c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Company.UpdateCompany(userId, companyId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) DeleteCompany(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	companyId, err := getId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Company.DeleteCompany(userId, companyId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type AllMultipliersResponse struct {
	Data []dto.MultiplierDto `json:"data"`
}

func (h *Handler) GetAllMultipliers(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	param := c.Query("sortBy")
	if param != "ebitda" && param != "pe" && param != "ev" {
		newErrorResponse(c, http.StatusBadRequest, "invalid sortBy param")
		return
	}

	multipliers, err := h.services.Company.GetAllMultipliers(userId, param)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, AllMultipliersResponse{
		Data: multipliers,
	})
}

func (h *Handler) GetMultiplierById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	companyId, err := getId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	mult, err := h.services.Company.GetMultiplierById(userId, companyId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, mult)
}
