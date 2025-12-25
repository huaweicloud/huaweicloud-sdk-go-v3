package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableDataclassTypeResponse Response Object
type EnableDataclassTypeResponse struct {

	// 成功响应码.
	Code *string `json:"code,omitempty"`

	// 成功信息
	Message *string `json:"message,omitempty"`

	// 成功标识.
	Data *string `json:"data,omitempty"`

	// 响应标识
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o EnableDataclassTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDataclassTypeResponse struct{}"
	}

	return strings.Join([]string{"EnableDataclassTypeResponse", string(data)}, " ")
}
