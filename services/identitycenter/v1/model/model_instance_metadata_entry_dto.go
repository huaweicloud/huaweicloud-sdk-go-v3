package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceMetadataEntryDto 提供关于IdentityCenter实例的信息
type InstanceMetadataEntryDto struct {

	// 关联到IAM身份中心的identity store的唯一标识
	IdentityStoreId string `json:"identity_store_id"`

	// IAM身份中心实例唯一标识.
	InstanceId string `json:"instance_id"`

	// 用户自定义的identity_store_id别名
	Alias *string `json:"alias,omitempty"`
}

func (o InstanceMetadataEntryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceMetadataEntryDto struct{}"
	}

	return strings.Join([]string{"InstanceMetadataEntryDto", string(data)}, " ")
}
