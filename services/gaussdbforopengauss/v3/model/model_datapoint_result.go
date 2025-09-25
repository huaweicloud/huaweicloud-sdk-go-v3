package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatapointResult struct {

	// **参数解释**: 指标项名，实例指标用实例ID、节点指标用节点名称、组件指标用组件名称。 **取值范围**: 不涉及。
	DatapointName string `json:"datapoint_name"`

	// **参数解释**: 指标值集合。
	DatapointValues []string `json:"datapoint_values"`
}

func (o DatapointResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatapointResult struct{}"
	}

	return strings.Join([]string{"DatapointResult", string(data)}, " ")
}
