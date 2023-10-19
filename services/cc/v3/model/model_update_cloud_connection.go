package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCloudConnection 更新云连接实例的详细信息。
type UpdateCloudConnection struct {

	// 实例名字。
	Name *string `json:"name,omitempty"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`
}

func (o UpdateCloudConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloudConnection struct{}"
	}

	return strings.Join([]string{"UpdateCloudConnection", string(data)}, " ")
}
