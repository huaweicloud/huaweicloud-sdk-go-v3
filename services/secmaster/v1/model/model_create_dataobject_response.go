package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataobjectResponse Response Object
type CreateDataobjectResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 是否成功
	Success *bool `json:"success,omitempty"`

	Data           *DataObjectCreateUpdateResponse `json:"data,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o CreateDataobjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataobjectResponse struct{}"
	}

	return strings.Join([]string{"CreateDataobjectResponse", string(data)}, " ")
}
