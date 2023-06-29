package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnlockWorkspaceRequestBody 解除云办公服务锁定状态请求。
type UnlockWorkspaceRequestBody struct {

	// 解除项目锁定操作类型。
	OperateType *string `json:"operate_type,omitempty"`
}

func (o UnlockWorkspaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockWorkspaceRequestBody struct{}"
	}

	return strings.Join([]string{"UnlockWorkspaceRequestBody", string(data)}, " ")
}
