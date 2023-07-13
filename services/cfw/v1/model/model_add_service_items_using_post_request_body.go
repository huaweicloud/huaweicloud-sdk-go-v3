package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddServiceItemsUsingPostRequestBody struct {

	// 服务组id
	SetId string `json:"set_id"`

	// 添加服务组成员
	ServiceItems []AddServiceItemsUsingPostRequestBodyServiceItems `json:"service_items"`
}

func (o AddServiceItemsUsingPostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServiceItemsUsingPostRequestBody struct{}"
	}

	return strings.Join([]string{"AddServiceItemsUsingPostRequestBody", string(data)}, " ")
}
