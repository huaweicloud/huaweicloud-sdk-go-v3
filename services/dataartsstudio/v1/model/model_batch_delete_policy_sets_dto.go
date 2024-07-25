package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeletePolicySetsDto struct {

	// 待删除动态脱敏策略id列表。
	Ids *[]string `json:"ids,omitempty"`
}

func (o BatchDeletePolicySetsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePolicySetsDto struct{}"
	}

	return strings.Join([]string{"BatchDeletePolicySetsDto", string(data)}, " ")
}
