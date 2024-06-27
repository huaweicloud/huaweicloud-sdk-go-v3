package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddServiceSetUsingPostRequestBody Create Service Set Required Body
type AddServiceSetUsingPostRequestBody struct {

	// 互联网边界防护对象id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，type为0的为互联网边界防护对象id。
	ObjectId string `json:"object_id"`

	// 服务组名称
	Name string `json:"name"`

	// 服务组描述信息
	Description *string `json:"description,omitempty"`
}

func (o AddServiceSetUsingPostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServiceSetUsingPostRequestBody struct{}"
	}

	return strings.Join([]string{"AddServiceSetUsingPostRequestBody", string(data)}, " ")
}
