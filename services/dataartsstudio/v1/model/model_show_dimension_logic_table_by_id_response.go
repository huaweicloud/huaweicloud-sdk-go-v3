package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDimensionLogicTableByIdResponse Response Object
type ShowDimensionLogicTableByIdResponse struct {
	Data           *ShowDimensionLogicTableByIdResultData `json:"data,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ShowDimensionLogicTableByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDimensionLogicTableByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowDimensionLogicTableByIdResponse", string(data)}, " ")
}
