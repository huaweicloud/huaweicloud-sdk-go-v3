package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAimMsgAppRequest Request Object
type UpdateAimMsgAppRequest struct {

	// 应用ID。
	AppId string `json:"app_id"`

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *MsgAppRequest `json:"body,omitempty"`
}

func (o UpdateAimMsgAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAimMsgAppRequest struct{}"
	}

	return strings.Join([]string{"UpdateAimMsgAppRequest", string(data)}, " ")
}
