package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddServiceItemsUsingPostRequestBody struct {

	// 服务组id，可通过[获取服务组列表接口](ListServiceSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。
	SetId string `json:"set_id"`

	// 服务组成员列表
	ServiceItems []AddServiceItemsUsingPostRequestBodyServiceItems `json:"service_items"`
}

func (o AddServiceItemsUsingPostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServiceItemsUsingPostRequestBody struct{}"
	}

	return strings.Join([]string{"AddServiceItemsUsingPostRequestBody", string(data)}, " ")
}
