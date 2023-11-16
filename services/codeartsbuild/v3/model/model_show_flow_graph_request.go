package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlowGraphRequest Request Object
type ShowFlowGraphRequest struct {

	// 父任务构建记录ID
	BuildFlowRecordId string `json:"build_flow_record_id"`
}

func (o ShowFlowGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowGraphRequest struct{}"
	}

	return strings.Join([]string{"ShowFlowGraphRequest", string(data)}, " ")
}
