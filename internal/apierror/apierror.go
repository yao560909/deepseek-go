package apierror

import (
	"fmt"
	"github.com/yao560909/deepseek-go/internal/apijson"
	"net/http"
)

type Error struct {
	Code       string    `json:"code,required,nullable"`
	Message    string    `json:"message,required"`
	Param      string    `json:"param,required,nullable"`
	Type       string    `json:"type,required"`
	JSON       errorJSON `json:"-"`
	StatusCode int
	Request    *http.Request
	Response   *http.Response
}

type errorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Param       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r errorJSON) RawJSON() string {
	return r.raw
}

func (r *Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r *Error) Error() string {
	// Attempt to re-populate the response body
	return fmt.Sprintf("%s \"%s\": %d %s %s", r.Request.Method, r.Request.URL, r.Response.StatusCode, http.StatusText(r.Response.StatusCode), r.JSON.RawJSON())
}
