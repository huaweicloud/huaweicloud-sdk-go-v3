package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupProtectedBranchesRequest Request Object
type ListGroupProtectedBranchesRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id **默认取值：** 不涉及。
	GroupId int32 `json:"group_id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 检索内容 **约束限制：** 保护分支名称 **取值范围：** 保护分支名称 **默认取值：** 不涉及
	Search *string `json:"search,omitempty"`

	// **参数解释：** 是否返回带有user_action结构的数据，user_action结构的数据为最新的结构，推荐传参为true **约束限制：** true,false **取值范围：** true,false **默认取值：** 默认不传参
	UserActions bool `json:"user_actions"`
}

func (o ListGroupProtectedBranchesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupProtectedBranchesRequest struct{}"
	}

	return strings.Join([]string{"ListGroupProtectedBranchesRequest", string(data)}, " ")
}
