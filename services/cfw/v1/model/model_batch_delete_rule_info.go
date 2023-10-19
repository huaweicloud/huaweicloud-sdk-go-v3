package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteRuleInfo struct {

	// acl名称
	Name *string `json:"name,omitempty"`

	// aclId
	Id *string `json:"id,omitempty"`
}

func (o BatchDeleteRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRuleInfo struct{}"
	}

	return strings.Join([]string{"BatchDeleteRuleInfo", string(data)}, " ")
}
