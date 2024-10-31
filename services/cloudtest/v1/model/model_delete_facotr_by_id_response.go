package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFacotrByIdResponse Response Object
type DeleteFacotrByIdResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteFacotrByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFacotrByIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteFacotrByIdResponse", string(data)}, " ")
}
