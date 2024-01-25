package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVolumeResponse Response Object
type ListVolumeResponse struct {

	// 硬盘数量。
	Count *int32 `json:"count,omitempty"`

	// 硬盘列表。
	Volumes        *[]Volume `json:"volumes,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumeResponse struct{}"
	}

	return strings.Join([]string{"ListVolumeResponse", string(data)}, " ")
}
