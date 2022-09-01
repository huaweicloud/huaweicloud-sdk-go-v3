package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowResJobRequest struct {

	// 工作空间id
	WorkspaceId string `json:"workspace_id" xml:"workspace_id"`

	// 资源id（数据源id 或 场景id）
	ResourceId string `json:"resource_id" xml:"resource_id"`

	// 类别： - RECALL，召回作业 - DATASOURCE，数据源作业 - FILTER，过滤作业 - SORTING，排序作业 - EVALUATE，效果评估作业
	Category string `json:"category" xml:"category"`
}

func (o ShowResJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResJobRequest struct{}"
	}

	return strings.Join([]string{"ShowResJobRequest", string(data)}, " ")
}
