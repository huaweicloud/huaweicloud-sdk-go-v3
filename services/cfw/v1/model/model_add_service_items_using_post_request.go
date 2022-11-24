package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddServiceItemsUsingPostRequest struct {
	Body *AddServiceItemsUsingPostRequestBody `json:"body,omitempty"`
}

func (o AddServiceItemsUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServiceItemsUsingPostRequest struct{}"
	}

	return strings.Join([]string{"AddServiceItemsUsingPostRequest", string(data)}, " ")
}
