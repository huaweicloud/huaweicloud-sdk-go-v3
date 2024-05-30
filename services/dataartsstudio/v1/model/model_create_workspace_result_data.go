package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkspaceResultData data，统一的返回结果的最外层数据结构。
type CreateWorkspaceResultData struct {
	Value *WorkspaceVo `json:"value,omitempty"`
}

func (o CreateWorkspaceResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkspaceResultData struct{}"
	}

	return strings.Join([]string{"CreateWorkspaceResultData", string(data)}, " ")
}
