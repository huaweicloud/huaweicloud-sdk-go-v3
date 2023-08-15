package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddServiceItemsUsingPostResponse Response Object
type AddServiceItemsUsingPostResponse struct {
	Data           *ServiceItemIds `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o AddServiceItemsUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServiceItemsUsingPostResponse struct{}"
	}

	return strings.Join([]string{"AddServiceItemsUsingPostResponse", string(data)}, " ")
}
