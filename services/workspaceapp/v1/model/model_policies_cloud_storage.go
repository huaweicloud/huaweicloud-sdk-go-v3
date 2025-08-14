package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesCloudStorage 云存储配置。
type PoliciesCloudStorage struct {

	// 云存储配置开关： false: 表示关闭 true: 表示开启
	CloudStorageEnable *bool `json:"cloud_storage_enable,omitempty"`

	Options *CloudStorageOptions `json:"options,omitempty"`
}

func (o PoliciesCloudStorage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesCloudStorage struct{}"
	}

	return strings.Join([]string{"PoliciesCloudStorage", string(data)}, " ")
}
