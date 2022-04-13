package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticStatusV2 struct {
	// 未解决

	Unresolved *int32 `json:"unresolved,omitempty"`
	// 已解决

	Resolved *int32 `json:"resolved,omitempty"`
	// 已忽略

	Dismissed *int32 `json:"dismissed,omitempty"`
}

func (o StatisticStatusV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticStatusV2 struct{}"
	}

	return strings.Join([]string{"StatisticStatusV2", string(data)}, " ")
}
