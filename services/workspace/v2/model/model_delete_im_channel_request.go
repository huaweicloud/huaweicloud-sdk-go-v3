package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImChannelRequest Request Object
type DeleteImChannelRequest struct {

	// Agent 实例主键 ID
	Id string `json:"id"`

	// IM 平台类型：wecom / feishu / dingtalk-connector
	Platform string `json:"platform"`
}

func (o DeleteImChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImChannelRequest struct{}"
	}

	return strings.Join([]string{"DeleteImChannelRequest", string(data)}, " ")
}
