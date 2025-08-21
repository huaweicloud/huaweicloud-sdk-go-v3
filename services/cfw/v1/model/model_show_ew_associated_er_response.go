package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEwAssociatedErResponse Response Object
type ShowEwAssociatedErResponse struct {
	Data           *ShowEwAssociatedErRespData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowEwAssociatedErResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEwAssociatedErResponse struct{}"
	}

	return strings.Join([]string{"ShowEwAssociatedErResponse", string(data)}, " ")
}
