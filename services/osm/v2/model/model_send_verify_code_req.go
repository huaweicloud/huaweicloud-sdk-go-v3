package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SendVerifyCodeReq struct {

	// 联系方式的值
	ContactValue string `json:"contact_value"`

	// 联系方式的类型，0：短信；1：邮件
	ContactWay int32 `json:"contact_way"`

	// 国家码
	AreaCode string `json:"area_code"`
}

func (o SendVerifyCodeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendVerifyCodeReq struct{}"
	}

	return strings.Join([]string{"SendVerifyCodeReq", string(data)}, " ")
}
