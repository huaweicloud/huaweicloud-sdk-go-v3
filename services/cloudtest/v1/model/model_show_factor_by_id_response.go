package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactorByIdResponse Response Object
type ShowFactorByIdResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFactorByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactorByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowFactorByIdResponse", string(data)}, " ")
}
