package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachmentInstanceId 中心网络附件对端实例ID，企业路由器的路由表ID或者GDGW的ID。
type AttachmentInstanceId struct {

	// 中心网络附件对端实例ID，企业路由器的路由表ID或者GDGW的ID。
	AttachmentInstanceId string `json:"attachment_instance_id"`
}

func (o AttachmentInstanceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentInstanceId struct{}"
	}

	return strings.Join([]string{"AttachmentInstanceId", string(data)}, " ")
}
