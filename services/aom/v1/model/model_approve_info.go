package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 审批信息
type ApproveInfo struct {

	// 审批人主题选择，当是否审核,默认是审核时，不能为空。
	TopicSelected *string `json:"topic_selected,omitempty"`

	// 是否审核,默认是不审核，true，false。
	NeedApprove *bool `json:"need_approve,omitempty"`

	// 审批主题urn集合。
	SmnUrnList *string `json:"smn_urn_list,omitempty"`
}

func (o ApproveInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApproveInfo struct{}"
	}

	return strings.Join([]string{"ApproveInfo", string(data)}, " ")
}
