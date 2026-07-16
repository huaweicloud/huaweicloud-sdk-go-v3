package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OutputModel **参数解释**：自定义训练作业产物输出信息。
type OutputModel struct {
	Obs *ObsModel `json:"obs,omitempty"`
}

func (o OutputModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputModel struct{}"
	}

	return strings.Join([]string{"OutputModel", string(data)}, " ")
}
