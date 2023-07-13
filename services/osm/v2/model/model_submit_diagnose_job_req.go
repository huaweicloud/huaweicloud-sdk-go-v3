package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubmitDiagnoseJobReq struct {

	// 任务类型，例如 ecs诊断任务 1，rds诊断任务 2
	Type int32 `json:"type"`

	// 类型对应的特有参数，例如ecs需要传eip,rds 需要传输instanceId
	Params map[string]string `json:"params,omitempty"`

	// 节点id
	RegionId *string `json:"region_id,omitempty"`
}

func (o SubmitDiagnoseJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubmitDiagnoseJobReq struct{}"
	}

	return strings.Join([]string{"SubmitDiagnoseJobReq", string(data)}, " ")
}
