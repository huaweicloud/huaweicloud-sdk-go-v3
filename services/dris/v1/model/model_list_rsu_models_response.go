package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRsuModelsResponse Response Object
type ListRsuModelsResponse struct {

	// **参数说明**：满足查询条件的记录总数。
	Count *int64 `json:"count,omitempty"`

	// **参数说明**：RSU型号信息列表。
	RsuModels      *[]RsuModelSummary `json:"rsu_models,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListRsuModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRsuModelsResponse struct{}"
	}

	return strings.Join([]string{"ListRsuModelsResponse", string(data)}, " ")
}
