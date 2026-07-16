package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateWorkspaceResponseBodyGrants struct {

	// 用户ID,此参数与user_name必填一个。两者都填优先使用user_id。
	UserId *string `json:"user_id,omitempty"`

	// IAM用户名称。此参数与user_id必填一个。
	UserName *string `json:"user_name,omitempty"`
}

func (o CreateWorkspaceResponseBodyGrants) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkspaceResponseBodyGrants struct{}"
	}

	return strings.Join([]string{"CreateWorkspaceResponseBodyGrants", string(data)}, " ")
}
