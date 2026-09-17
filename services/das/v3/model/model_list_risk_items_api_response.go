package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRiskItemsApiResponse Response Object
type ListRiskItemsApiResponse struct {

	// 数据库类型
	EngineType *string `json:"engine_type,omitempty"`

	// 指标阈值列表
	Items          *[]RiskItemInfo `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListRiskItemsApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskItemsApiResponse struct{}"
	}

	return strings.Join([]string{"ListRiskItemsApiResponse", string(data)}, " ")
}
