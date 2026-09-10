package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteWorkspacesRequestBody 删除工作空间请求体
type BatchDeleteWorkspacesRequestBody struct {

	// 待删除的工作空间ID列表，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	WorkspaceIds []string `json:"workspace_ids"`
}

func (o BatchDeleteWorkspacesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteWorkspacesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteWorkspacesRequestBody", string(data)}, " ")
}
