package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataobjectResponse Response Object
type UpdateDataobjectResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 是否成功
	Success *bool `json:"success,omitempty"`

	Data           *DataObjectDetail `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateDataobjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataobjectResponse struct{}"
	}

	return strings.Join([]string{"UpdateDataobjectResponse", string(data)}, " ")
}
