package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupSumDto struct {

	// 打开状态MR计数
	OpenMergeRequestsCount *int32 `json:"open_merge_requests_count,omitempty"`
}

func (o GroupSumDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupSumDto struct{}"
	}

	return strings.Join([]string{"GroupSumDto", string(data)}, " ")
}
