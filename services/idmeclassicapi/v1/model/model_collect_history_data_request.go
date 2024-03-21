package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectHistoryDataRequest Request Object
type CollectHistoryDataRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoStatisticsPvo `json:"body,omitempty"`
}

func (o CollectHistoryDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectHistoryDataRequest struct{}"
	}

	return strings.Join([]string{"CollectHistoryDataRequest", string(data)}, " ")
}
