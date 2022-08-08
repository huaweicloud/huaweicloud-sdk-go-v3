package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新云连接实例的详细信息。
type UpdateCloudConnection struct {

	// 云连接实例的名字。只能由中文、英文字母、数字、下划线、中划线、点组成。
	Name *string `json:"name,omitempty"`

	// 云连接实例的描述。不支持 <>。
	Description *string `json:"description,omitempty"`
}

func (o UpdateCloudConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloudConnection struct{}"
	}

	return strings.Join([]string{"UpdateCloudConnection", string(data)}, " ")
}
