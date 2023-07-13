package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVolumeTypeResponse Response Object
type ListVolumeTypeResponse struct {

	// 磁盘列表
	VolumeTypes    *[]VolumeTypeInfo `json:"volume_types,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListVolumeTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumeTypeResponse struct{}"
	}

	return strings.Join([]string{"ListVolumeTypeResponse", string(data)}, " ")
}
