package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopologyLabels struct {

	// **参数解释**：转发策略信息。
	L7policies *[]PolicyLabel `json:"l7policies,omitempty"`
}

func (o TopologyLabels) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopologyLabels struct{}"
	}

	return strings.Join([]string{"TopologyLabels", string(data)}, " ")
}
