package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteRuleInfo struct {

	// 批量删除的acl的名称
	Name *string `json:"name,omitempty"`

	// 批量删除的acl的id
	Id *string `json:"id,omitempty"`
}

func (o BatchDeleteRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRuleInfo struct{}"
	}

	return strings.Join([]string{"BatchDeleteRuleInfo", string(data)}, " ")
}
