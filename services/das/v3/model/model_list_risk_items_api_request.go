package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRiskItemsApiRequest Request Object
type ListRiskItemsApiRequest struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`
}

func (o ListRiskItemsApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskItemsApiRequest struct{}"
	}

	return strings.Join([]string{"ListRiskItemsApiRequest", string(data)}, " ")
}
