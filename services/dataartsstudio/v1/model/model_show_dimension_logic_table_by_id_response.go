package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDimensionLogicTableByIdResponse Response Object
type ShowDimensionLogicTableByIdResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowDimensionLogicTableByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDimensionLogicTableByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowDimensionLogicTableByIdResponse", string(data)}, " ")
}
