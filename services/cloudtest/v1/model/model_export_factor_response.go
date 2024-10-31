package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFactorResponse Response Object
type ExportFactorResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportFactorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFactorResponse struct{}"
	}

	return strings.Join([]string{"ExportFactorResponse", string(data)}, " ")
}
