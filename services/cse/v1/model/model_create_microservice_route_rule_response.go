package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMicroserviceRouteRuleResponse Response Object
type CreateMicroserviceRouteRuleResponse struct {

	// 结果信息
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMicroserviceRouteRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMicroserviceRouteRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateMicroserviceRouteRuleResponse", string(data)}, " ")
}
