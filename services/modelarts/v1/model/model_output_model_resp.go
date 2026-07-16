package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OutputModelResp **参数解释**：自定义训练作业产物输出信息。
type OutputModelResp struct {
	Obs *ObsModelResp `json:"obs,omitempty"`
}

func (o OutputModelResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputModelResp struct{}"
	}

	return strings.Join([]string{"OutputModelResp", string(data)}, " ")
}
