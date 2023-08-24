package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimMsgAppDetailRequest Request Object
type ListAimMsgAppDetailRequest struct {

	// 应用ID。
	AppId string `json:"app_id"`

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`
}

func (o ListAimMsgAppDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimMsgAppDetailRequest struct{}"
	}

	return strings.Join([]string{"ListAimMsgAppDetailRequest", string(data)}, " ")
}
