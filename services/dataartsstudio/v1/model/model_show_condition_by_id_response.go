package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConditionByIdResponse Response Object
type ShowConditionByIdResponse struct {
	Data           *ShowConditionByIdResultData `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowConditionByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConditionByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowConditionByIdResponse", string(data)}, " ")
}
