package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShareTypeUsageCapacity struct {

	// 总容量
	Total *int32 `json:"total,omitempty"`

	// 已用容量
	Usage *int32 `json:"usage,omitempty"`
}

func (o ShareTypeUsageCapacity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareTypeUsageCapacity struct{}"
	}

	return strings.Join([]string{"ShareTypeUsageCapacity", string(data)}, " ")
}
