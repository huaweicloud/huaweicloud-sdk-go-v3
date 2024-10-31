package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServiceItemResponse Response Object
type DeleteServiceItemResponse struct {
	Data           *DeleteServiceItemResponseBodyData `json:"data,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o DeleteServiceItemResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceItemResponse struct{}"
	}

	return strings.Join([]string{"DeleteServiceItemResponse", string(data)}, " ")
}
