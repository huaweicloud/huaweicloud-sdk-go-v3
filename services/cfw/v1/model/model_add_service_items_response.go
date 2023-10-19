package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddServiceItemsResponse Response Object
type AddServiceItemsResponse struct {
	Data           *ServiceItemIds `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o AddServiceItemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServiceItemsResponse struct{}"
	}

	return strings.Join([]string{"AddServiceItemsResponse", string(data)}, " ")
}
