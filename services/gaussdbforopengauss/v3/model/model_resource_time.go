package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTime 资源耗时信息
type ResourceTime struct {

	// **参数解释**: 总耗时（单位：微秒）。 **取值范围**: 不涉及。
	AllTime int64 `json:"all_time"`

	ResourceTimeDetails *ResourceTimeDetails `json:"resource_time_details"`
}

func (o ResourceTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTime struct{}"
	}

	return strings.Join([]string{"ResourceTime", string(data)}, " ")
}
