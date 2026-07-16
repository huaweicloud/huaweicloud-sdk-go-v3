package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ViewWorkspaceResponseGrants struct {

	// IAM用户ID。此参数与user_name必填一个。两者都填优先使用user_id。
	UserId *string `json:"user_id,omitempty"`

	// IAM用户名称。此参数与user_id必填一个。
	UserName *string `json:"user_name,omitempty"`

	// 参数解释： 授权用户类型。 约束限制： 如果是联邦用户或者委托用户的话必填。 取值范围： IAM:IAM用户, FEDERATE：联邦用户, AGENCY：委托用户。 默认取值： IAM。
	UserType *string `json:"user_type,omitempty"`
}

func (o ViewWorkspaceResponseGrants) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ViewWorkspaceResponseGrants struct{}"
	}

	return strings.Join([]string{"ViewWorkspaceResponseGrants", string(data)}, " ")
}
