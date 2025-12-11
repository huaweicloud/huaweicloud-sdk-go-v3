package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociationRoutePolicyRequestBody struct {
	RoutePolicy *ExportRoutePolicy `json:"route_policy"`
}

func (o AssociationRoutePolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociationRoutePolicyRequestBody struct{}"
	}

	return strings.Join([]string{"AssociationRoutePolicyRequestBody", string(data)}, " ")
}
