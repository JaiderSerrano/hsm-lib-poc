package handler

import (
	"net/http"
	"strconv"

	"github.com/mercadolibre/fury_go-core/pkg/log"
	"github.com/mercadolibre/fury_go-core/pkg/web"
	"github.com/mercadolibre/hsm-lib-poc/internal/hsm"
)

type handler struct {
	service hsm.Service
}

func NewHSMHandler(service hsm.Service) *handler {
	return &handler{service}
}

type response struct {
	Message string `json:"message"`
}

func (h handler) ARQCValidation(w http.ResponseWriter, r *http.Request) error {
	var arqcParams hsm.ARQCParams
	err := web.DecodeJSON(r, &arqcParams)

	if err != nil {
		log.Error(r.Context(), "error decoding json ARQC parameters", log.Err(err))
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.service.ARQCValidation(r.Context(), arqcParams)
	if err != nil {
		log.Error(r.Context(), "error validating ARQC.", log.Err(err))
		return web.NewError(http.StatusInternalServerError, err.Error())
	}

	res := response{Message: strconv.FormatBool(resp)}
	return web.EncodeJSON(w, res, http.StatusOK)
}

func (h handler) PINGeneration(w http.ResponseWriter, r *http.Request) error {
	var pinGenParams hsm.PINGenerationParams
	err := web.DecodeJSON(r, &pinGenParams)
	if err != nil {
		log.Error(r.Context(), "error decoding json PIN generation parameters", log.Err(err))
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.service.PINGeneration(r.Context(), pinGenParams)
	if err != nil {
		log.Error(r.Context(), "error generating PIN.", log.Err(err))
		return web.NewError(http.StatusInternalServerError, err.Error())
	}

	return web.EncodeJSON(w, resp, http.StatusOK)
}

func (h handler) PVVGeneration(w http.ResponseWriter, r *http.Request) error {
	var pvvGenParams hsm.PVVGenerationParams
	err := web.DecodeJSON(r, &pvvGenParams)
	if err != nil {
		log.Error(r.Context(), "error decoding json PVV generation parameters", log.Err(err))
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.service.PVVGeneration(r.Context(), pvvGenParams)
	if err != nil {
		log.Error(r.Context(), "error generating PIN.", log.Err(err))
		return web.NewError(http.StatusInternalServerError, err.Error())
	}

	return web.EncodeJSON(w, resp, http.StatusOK)
}
