package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateApplicationViewRequestBodySyncRules struct {

	// 企业项目id
	EpId *string `json:"ep_id,omitempty"`

	// 关联标签
	RuleTags *string `json:"rule_tags,omitempty"`
}

func (o BatchCreateApplicationViewRequestBodySyncRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateApplicationViewRequestBodySyncRules struct{}"
	}

	return strings.Join([]string{"BatchCreateApplicationViewRequestBodySyncRules", string(data)}, " ")
}
