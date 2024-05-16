package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HtapStorageTypeStorageType struct {

	// 存储介质名。
	Name *string `json:"name,omitempty"`

	// 可支持可用区信息。
	AzStatus map[string]string `json:"az_status,omitempty"`

	// 最小可提供磁盘大小。
	MinVolumeSize *int32 `json:"min_volume_size,omitempty"`
}

func (o HtapStorageTypeStorageType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapStorageTypeStorageType struct{}"
	}

	return strings.Join([]string{"HtapStorageTypeStorageType", string(data)}, " ")
}
