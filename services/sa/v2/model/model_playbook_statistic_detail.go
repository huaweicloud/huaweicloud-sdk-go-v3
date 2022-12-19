package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// statistics of playbook
type PlaybookStatisticDetail struct {

	// 未审核剧本数量
	UnapprovedNum *int32 `json:"unapproved_num,omitempty"`

	// 未启用剧本数量
	DisabledNum *int32 `json:"disabled_num,omitempty"`

	// 已启用剧本数量
	EnabledNum *int32 `json:"enabled_num,omitempty"`
}

func (o PlaybookStatisticDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlaybookStatisticDetail struct{}"
	}

	return strings.Join([]string{"PlaybookStatisticDetail", string(data)}, " ")
}
