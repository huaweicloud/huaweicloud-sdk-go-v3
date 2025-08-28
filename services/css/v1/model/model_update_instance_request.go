package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceRequest Request Object
type UpdateInstanceRequest struct {

	// 指定替换集群ID。
	ClusterId string `json:"cluster_id"`

	// 指定替换节点ID。
	InstanceId string `json:"instance_id"`

	// 是否迁移数据。
	MigrateData *string `json:"migrateData,omitempty"`

	// 委托名称，委托给CSS服务，允许CSS调用您的其他云服务。
	Agency *string `json:"agency,omitempty"`
}

func (o UpdateInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRequest", string(data)}, " ")
}
