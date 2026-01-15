package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesCloudStorage 云存储。
type PoliciesCloudStorage struct {

	// 是否开启云存储。取值为： false：表示关闭。 true：表示开启。
	CloudStorageEnable *bool `json:"cloud_storage_enable,omitempty"`

	Options *PoliciesCloudStorageOptions `json:"options,omitempty"`
}

func (o PoliciesCloudStorage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesCloudStorage struct{}"
	}

	return strings.Join([]string{"PoliciesCloudStorage", string(data)}, " ")
}
