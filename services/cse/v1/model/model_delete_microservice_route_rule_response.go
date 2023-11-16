package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMicroserviceRouteRuleResponse Response Object
type DeleteMicroserviceRouteRuleResponse struct {

	// 结果信息
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMicroserviceRouteRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMicroserviceRouteRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteMicroserviceRouteRuleResponse", string(data)}, " ")
}
