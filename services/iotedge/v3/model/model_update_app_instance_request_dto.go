package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAppInstanceRequestDto struct {

	// 动作类型
	Action string `json:"action"`

	// 命名空间，应用实例部署于非默认命名空间(default)时必填
	Namespace *string `json:"namespace,omitempty"`

	// 升级的目标版本号，动作类型为upgrade时必填
	AppVersion *string `json:"app_version,omitempty"`

	// 应用实例chart配置，动作类型为upgrade时必填
	Values *interface{} `json:"values,omitempty"`

	// 回退的目标版本号，动作类型为rollback时必填
	RollbackVersion *string `json:"rollback_version,omitempty"`
}

func (o UpdateAppInstanceRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppInstanceRequestDto struct{}"
	}

	return strings.Join([]string{"UpdateAppInstanceRequestDto", string(data)}, " ")
}
