package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyBaselinePolicyGroupRequestBody 被复制的基线策略组
type CopyBaselinePolicyGroupRequestBody struct {

	// 新策略组描述信息
	Description *string `json:"description,omitempty"`

	// 新策略组名称
	GroupName string `json:"group_name"`

	// 被复制策略组id
	GroupId string `json:"group_id"`
}

func (o CopyBaselinePolicyGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyBaselinePolicyGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CopyBaselinePolicyGroupRequestBody", string(data)}, " ")
}
