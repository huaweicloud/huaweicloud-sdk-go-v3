package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMicroserviceRouteRuleResponse Response Object
type ListMicroserviceRouteRuleResponse struct {

	// 结果信息
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMicroserviceRouteRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMicroserviceRouteRuleResponse struct{}"
	}

	return strings.Join([]string{"ListMicroserviceRouteRuleResponse", string(data)}, " ")
}
