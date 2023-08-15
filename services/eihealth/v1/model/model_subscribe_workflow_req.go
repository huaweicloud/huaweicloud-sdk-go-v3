package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeWorkflowReq 订阅流程请求体
type SubscribeWorkflowReq struct {

	// 资产id。长度1-64，只能包含字母、数字、下划线和中划线
	AssetId string `json:"asset_id"`

	// 资产版本。长度1-128，字母或数字开头，后面跟小写字母、数字、小数点、斜杠、下划线或中划线
	AssetVersion string `json:"asset_version"`

	// 目标流程名称。取值范围[1,56]，允许大小写字母、数字、以及特殊字符中划线(-)和下划线(_)。
	DestinationWorkflowName string `json:"destination_workflow_name"`

	// 目标流程版本。取值范围：长度[1,24]，以小写字母或数字或大写字母开头，允许出现中划线，必须以小写字母或数字或大写字母结尾。
	DestinationWorkflowVersion string `json:"destination_workflow_version"`
}

func (o SubscribeWorkflowReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeWorkflowReq struct{}"
	}

	return strings.Join([]string{"SubscribeWorkflowReq", string(data)}, " ")
}
