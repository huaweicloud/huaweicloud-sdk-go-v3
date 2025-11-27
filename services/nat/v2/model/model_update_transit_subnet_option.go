package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTransitSubnetOption 更新中转子网请求体
type UpdateTransitSubnetOption struct {

	// 中转子网的名字。仅支持数字、字母、_（下划线）、-（中划线）、中文。 只有中转子网的申请方才可修改中转子网的名字，审批方不可修改。
	Name *string `json:"name,omitempty"`

	// 中转子网的描述。长度范围小于等于255个字符，不能包含“<”和“>”。 只有中转子网的申请方才可修改中转子网的描述，审批方不可修改。
	Description *string `json:"description,omitempty"`
}

func (o UpdateTransitSubnetOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTransitSubnetOption struct{}"
	}

	return strings.Join([]string{"UpdateTransitSubnetOption", string(data)}, " ")
}
