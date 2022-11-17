package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type GetStackTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o GetStackTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetStackTemplateResponse struct{}"
	}

	return strings.Join([]string{"GetStackTemplateResponse", string(data)}, " ")
}
