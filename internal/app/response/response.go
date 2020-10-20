package response

import (
	"context"
	"encoding/json"
	"net/http"
	"pmhb-book-service/internal/app/utils"
	"pmhb-book-service/internal/kerrors"
	"pmhb-book-service/internal/pkg/klog"
	"time"
)

// SuccessResponseFormat struct contains success response data
type SuccessResponseFormat struct {
	ServiceResponseBody *interface{}
}

// WriteJSON writes JSON data into responseWriter
func WriteJSON(w http.ResponseWriter) func(resp interface{}, httpStatusCode int) {
	return func(resp interface{}, httpStatusCode int) {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(httpStatusCode)
	}
}

// HandleError function returns failure response
func HandleError(r *http.Request, err error) (interface{}, int) {
	ctx := r.Context()
	logger := klog.WithPrefix("response")
	var res utils.KbankResponseHeader

	switch v := err.(type) {
	case kerrors.KError:
		logger.WithFields(v.Extract()).KError(ctx, v.LogMessage)

		errCode, errMsg, originalErrorStr := v.Code.String(), v.Message.String(), v.OriginalErrorStr
		res = utils.KbankResponseHeader{
			ResponseAppID: utils.ResponseAppID,
			ResponseDate:  time.Now(),
			StatusCode:    kerrors.StatusErrorFailed.String(),
			Errors: utils.ResponseErrorKbankHeader{
				ErrorCode:        errCode,
				ErrorDesc:        errMsg,
				OriginalErrorStr: originalErrorStr,
			},
		}
	default:
		logger.KError(ctx, v.Error())
	}
	c := context.WithValue(ctx, "request_status", kerrors.StatusErrorFailed.String())
	temp := r.WithContext(c)
	*r = *temp

	return res, http.StatusOK
}

// HandleSuccess function responses success response format
func HandleSuccess(r *http.Request, data interface{}) (interface{}, int) {
	rs := data
	c := context.WithValue(r.Context(), "request_status", kerrors.StatusNoneError.String())
	temp := r.WithContext(c)
	*r = *temp

	if r.Method == http.MethodPost {
		return rs, http.StatusCreated
	}
	return rs, http.StatusOK
}
