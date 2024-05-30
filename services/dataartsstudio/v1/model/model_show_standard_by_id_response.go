package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStandardByIdResponse Response Object
type ShowStandardByIdResponse struct {
	Data           *ShowStandardByIdResultData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowStandardByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStandardByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowStandardByIdResponse", string(data)}, " ")
}
