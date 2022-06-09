package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddOrModifyAttributeReq struct {

	// 自定义属性名称
	CustAttrName string `json:"cust_attr_name"`
}

func (o AddOrModifyAttributeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOrModifyAttributeReq struct{}"
	}

	return strings.Join([]string{"AddOrModifyAttributeReq", string(data)}, " ")
}
