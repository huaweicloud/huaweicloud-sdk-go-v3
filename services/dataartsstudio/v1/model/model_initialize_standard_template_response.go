package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InitializeStandardTemplateResponse Response Object
type InitializeStandardTemplateResponse struct {
	Data           *InitializeStandardTemplateResultData `json:"data,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o InitializeStandardTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InitializeStandardTemplateResponse struct{}"
	}

	return strings.Join([]string{"InitializeStandardTemplateResponse", string(data)}, " ")
}
