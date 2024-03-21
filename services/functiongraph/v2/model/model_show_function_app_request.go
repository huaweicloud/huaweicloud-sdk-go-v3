package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFunctionAppRequest Request Object
type ShowFunctionAppRequest struct {

	// 应用ID。
	Id string `json:"id"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ShowFunctionAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionAppRequest struct{}"
	}

	return strings.Join([]string{"ShowFunctionAppRequest", string(data)}, " ")
}
