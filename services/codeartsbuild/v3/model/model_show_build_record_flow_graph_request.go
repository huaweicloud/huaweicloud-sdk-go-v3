package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBuildRecordFlowGraphRequest Request Object
type ShowBuildRecordFlowGraphRequest struct {

	// 父任务构建记录ID
	BuildFlowRecordId string `json:"build_flow_record_id"`
}

func (o ShowBuildRecordFlowGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuildRecordFlowGraphRequest struct{}"
	}

	return strings.Join([]string{"ShowBuildRecordFlowGraphRequest", string(data)}, " ")
}
