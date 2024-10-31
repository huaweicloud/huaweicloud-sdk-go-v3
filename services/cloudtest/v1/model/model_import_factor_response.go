package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportFactorResponse Response Object
type ImportFactorResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportFactorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportFactorResponse struct{}"
	}

	return strings.Join([]string{"ImportFactorResponse", string(data)}, " ")
}
