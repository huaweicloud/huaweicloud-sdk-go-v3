package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlowTopResponse Response Object
type ShowFlowTopResponse struct {

	// **参数解释**： TOP统计信息 **取值范围**： 不涉及
	Data           map[string][]ItemVo `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowFlowTopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowTopResponse struct{}"
	}

	return strings.Join([]string{"ShowFlowTopResponse", string(data)}, " ")
}
