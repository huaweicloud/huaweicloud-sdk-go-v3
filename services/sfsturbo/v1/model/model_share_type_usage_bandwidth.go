package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShareTypeUsageBandwidth struct {

	// 总带宽
	Total *int32 `json:"total,omitempty"`

	// 已用带宽
	Usage *int32 `json:"usage,omitempty"`
}

func (o ShareTypeUsageBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareTypeUsageBandwidth struct{}"
	}

	return strings.Join([]string{"ShareTypeUsageBandwidth", string(data)}, " ")
}
