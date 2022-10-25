package handler

import (
	"fmt"
	"net/http"

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

func (h handler) ARQCValidation(w http.ResponseWriter, r *http.Request) error {
	var arqcParams hsm.ARQCParams
	err := web.DecodeJSON(r, &arqcParams)

	if err != nil {
		log.Error(r.Context(), "error decoding json ARQC parameters", log.Err(err))
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.service.ARQCValidation(r.Context(), arqcParams)
	if err != nil {
		fmt.Println(err.Error())
		log.Error(r.Context(), "error validating ARQC.", log.Err(err))
		return web.NewError(http.StatusInternalServerError, err.Error())
	}

	return web.EncodeJSON(w, resp, http.StatusOK)
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
		log.Error(r.Context(), "error generating PIN and PVV.", log.Err(err))
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
		log.Error(r.Context(), "error generating PVV.", log.Err(err))
		return web.NewError(http.StatusInternalServerError, err.Error())
	}

	return web.EncodeJSON(w, resp, http.StatusOK)
}

func (h handler) PINBlockGeneration(w http.ResponseWriter, r *http.Request) error {
	var pbGenParams hsm.PINBlockGenerationParams
	err := web.DecodeJSON(r, &pbGenParams)
	if err != nil {
		log.Error(r.Context(), "error decoding json PIN block generation parameters", log.Err(err))
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.service.PINBlockGeneration(r.Context(), pbGenParams)
	if err != nil {
		log.Error(r.Context(), "error generating PIN block.", log.Err(err))
		return web.NewError(http.StatusInternalServerError, err.Error())
	}

	return web.EncodeJSON(w, resp, http.StatusOK)
}

func (h handler) PINVerification(w http.ResponseWriter, r *http.Request) error {
	var pinVerParams hsm.PINVerificationParams
	err := web.DecodeJSON(r, &pinVerParams)
	if err != nil {
		log.Error(r.Context(), "error decoding json PIN validation parameters", log.Err(err))
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.service.PINValidation(r.Context(), pinVerParams)
	if err != nil {
		log.Error(r.Context(), "error generating PIN validation.", log.Err(err))
		return web.NewError(http.StatusInternalServerError, err.Error())
	}

	return web.EncodeJSON(w, resp, http.StatusOK)
}

func (h handler) ARPCGeneration(w http.ResponseWriter, r *http.Request) error {
	var arpcParams hsm.ARPCParams
	err := web.DecodeJSON(r, &arpcParams)
	if err != nil {
		log.Error(r.Context(), "error decoding json ARPC parameters", log.Err(err))
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	arpc, err := h.service.ARPCGeneration(r.Context(), arpcParams)
	if err != nil {
		log.Error(r.Context(), "error generating ARPC Ccryptogram", log.Err(err))
		return web.NewError(http.StatusInternalServerError, err.Error())
	}

	return web.EncodeJSON(w, arpc, http.StatusOK)
}

func (h handler) GenerateValidationData(w http.ResponseWriter, r *http.Request) error {
	var verificationData hsm.VerificationData

	err := web.DecodeJSON(r, &verificationData)
	if err != nil {
		log.Error(r.Context(), "error decoding json verification data parameters", log.Err(err))
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	vd, err := h.service.GenerateVerificationData(r.Context(), verificationData)
	if err != nil {
		log.Error(r.Context(), "error generating verification data", log.Err(err))
		return web.NewError(http.StatusInternalServerError, err.Error())
	}

	return web.EncodeJSON(w, vd, http.StatusOK)
}

func (h handler) ValidateValidationData(w http.ResponseWriter, r *http.Request) error {
	var verificationData hsm.VerificationData

	err := web.DecodeJSON(r, &verificationData)
	if err != nil {
		log.Error(r.Context(), "error decoding json verification data parameters", log.Err(err))
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	vd, err := h.service.ValidateVerificationData(r.Context(), verificationData)
	if err != nil {
		log.Error(r.Context(), "error validating verification data", log.Err(err))
		return web.NewError(http.StatusInternalServerError, err.Error())
	}

	return web.EncodeJSON(w, vd, http.StatusOK)
}
