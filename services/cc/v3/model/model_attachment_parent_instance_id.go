package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachmentParentInstanceId 连接的父资源ID，这里表示企业路由器ID。
type AttachmentParentInstanceId struct {

	// 连接的父资源ID，这里表示企业路由器ID。
	AttachmentParentInstanceId *string `json:"attachment_parent_instance_id,omitempty"`
}

func (o AttachmentParentInstanceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentParentInstanceId struct{}"
	}

	return strings.Join([]string{"AttachmentParentInstanceId", string(data)}, " ")
}
