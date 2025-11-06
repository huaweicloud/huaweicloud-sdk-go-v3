package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeRequestReviewerExternalDto struct {

	// 评审人id
	Id *int32 `json:"id,omitempty"`

	// 评审人名称
	Username *string `json:"username,omitempty"`

	// 评审人名称
	Name *string `json:"name,omitempty"`

	// 评审人昵称
	NickName *string `json:"nick_name,omitempty"`

	// 评审人状态
	State *bool `json:"state,omitempty"`

	// 打分
	Score *int32 `json:"score,omitempty"`

	// 租户名称
	TenantName *string `json:"tenant_name,omitempty"`
}

func (o MergeRequestReviewerExternalDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestReviewerExternalDto struct{}"
	}

	return strings.Join([]string{"MergeRequestReviewerExternalDto", string(data)}, " ")
}
