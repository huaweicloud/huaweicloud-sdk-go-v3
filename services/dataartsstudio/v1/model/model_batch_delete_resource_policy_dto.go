package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteResourcePolicyDto struct {

	// 待删除资源权限策略id列表。
	Ids *[]string `json:"ids,omitempty"`
}

func (o BatchDeleteResourcePolicyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourcePolicyDto struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourcePolicyDto", string(data)}, " ")
}
