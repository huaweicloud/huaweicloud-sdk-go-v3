package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachmentId 中心网络附件对端实例的连接ID，企业路由器的连接ID或者GDGW的连接ID。
type AttachmentId struct {

	// 中心网络附件对端实例的连接ID，企业路由器的连接ID或者GDGW的连接ID。
	AttachmentId *string `json:"attachment_id,omitempty"`
}

func (o AttachmentId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentId struct{}"
	}

	return strings.Join([]string{"AttachmentId", string(data)}, " ")
}
