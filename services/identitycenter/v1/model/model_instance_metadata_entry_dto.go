package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceMetadataEntryDto 提供关于IAM身份中心实例的信息
type InstanceMetadataEntryDto struct {

	// 关联到IAM身份中心实例的身份源的全局唯一标识符（ID）
	IdentityStoreId string `json:"identity_store_id"`

	// IAM身份中心实例的全局唯一标识符（ID）
	InstanceId string `json:"instance_id"`

	// 用户为身份源标识符定义的别名
	Alias *string `json:"alias,omitempty"`

	// 实例的统一资源名称（URN）
	InstanceUrn *string `json:"instance_urn,omitempty"`
}

func (o InstanceMetadataEntryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceMetadataEntryDto struct{}"
	}

	return strings.Join([]string{"InstanceMetadataEntryDto", string(data)}, " ")
}
