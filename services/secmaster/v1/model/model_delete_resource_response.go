package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceResponse Response Object
type DeleteResourceResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data           *BatchOperateResourceResult `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o DeleteResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteResourceResponse", string(data)}, " ")
}
