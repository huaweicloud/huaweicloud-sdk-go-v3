package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLogstashTemplateResponse Response Object
type DeleteLogstashTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLogstashTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogstashTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteLogstashTemplateResponse", string(data)}, " ")
}
