package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopUrl struct {

	// 总播放次数。
	Value *int64 `json:"value,omitempty" xml:"value"`

	// 媒资ID。
	AssetId *string `json:"asset_id,omitempty" xml:"asset_id"`

	// 媒资名称。
	Title *string `json:"title,omitempty" xml:"title"`

	// 媒资时长。  单位：秒。
	Duration *int32 `json:"duration,omitempty" xml:"duration"`

	// 媒资原始大小。  单位：字节。
	Size *int64 `json:"size,omitempty" xml:"size"`
}

func (o TopUrl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopUrl struct{}"
	}

	return strings.Join([]string{"TopUrl", string(data)}, " ")
}
