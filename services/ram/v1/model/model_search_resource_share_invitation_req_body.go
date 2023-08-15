package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchResourceShareInvitationReqBody struct {

	// 资源共享实例的ID列表。
	ResourceShareIds *[]string `json:"resource_share_ids,omitempty"`

	// 资源共享邀请的ID列表。
	ResourceShareInvitationIds *[]string `json:"resource_share_invitation_ids,omitempty"`

	// 资源共享邀请的当前状态。
	Status *string `json:"status,omitempty"`

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 分页位置标识。从marker指定索引的下一条数据开始查询。查询第一页数据时，不需要传入此参数，查询后续页码数据时，将查询前一页数据响应体中marker值配入此参数。
	Marker *string `json:"marker,omitempty"`
}

func (o SearchResourceShareInvitationReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResourceShareInvitationReqBody struct{}"
	}

	return strings.Join([]string{"SearchResourceShareInvitationReqBody", string(data)}, " ")
}
