package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactLogicTableByIdResponse Response Object
type ShowFactLogicTableByIdResponse struct {
	Data           *ShowFactLogicTableByIdResultData `json:"data,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ShowFactLogicTableByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactLogicTableByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowFactLogicTableByIdResponse", string(data)}, " ")
}
