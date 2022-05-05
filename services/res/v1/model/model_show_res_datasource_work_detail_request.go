package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowResDatasourceWorkDetailRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	// 资源id。
	ResourceId string `json:"resource_id"`

	// 任务类型： - DATA_STRUCT，数据结构 - DATA_INSPECTION，数据检测 - DATA_EXPLORATION，数据探索
	Type string `json:"type"`
}

func (o ShowResDatasourceWorkDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResDatasourceWorkDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowResDatasourceWorkDetailRequest", string(data)}, " ")
}
