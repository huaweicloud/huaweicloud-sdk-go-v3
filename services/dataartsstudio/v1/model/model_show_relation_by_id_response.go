package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRelationByIdResponse Response Object
type ShowRelationByIdResponse struct {
	Data           *ShowRelationByIdResultData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowRelationByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRelationByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowRelationByIdResponse", string(data)}, " ")
}
