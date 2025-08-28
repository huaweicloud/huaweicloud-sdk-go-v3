package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloneLoadbalancerResponseBodyLoadbalancerList struct {

	// **参数解释**：新实例的ID。  **取值范围**：不涉及
	Id *string `json:"id,omitempty"`
}

func (o CloneLoadbalancerResponseBodyLoadbalancerList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneLoadbalancerResponseBodyLoadbalancerList struct{}"
	}

	return strings.Join([]string{"CloneLoadbalancerResponseBodyLoadbalancerList", string(data)}, " ")
}
