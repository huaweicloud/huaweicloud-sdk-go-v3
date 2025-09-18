package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRedistributionControlResponse Response Object
type UpdateRedistributionControlResponse struct {

	// **参数解释**: 修改重分布参数的任务ID。 **取值范围**: 不涉及。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateRedistributionControlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRedistributionControlResponse struct{}"
	}

	return strings.Join([]string{"UpdateRedistributionControlResponse", string(data)}, " ")
}
