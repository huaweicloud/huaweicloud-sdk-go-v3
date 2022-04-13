package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VisionActiveCodeDto struct {
	// 终端类型 - idea-hub:智能协作大屏 - huawei-vision:智慧屏TV - welink-desktop(iwb):电子白板

	DevType string `json:"devType"`
	// 部门编号，若不携带则默认根部门

	DeptCode *string `json:"deptCode,omitempty"`
	// 终端的名称

	DevName string `json:"devName"`
	// 描述信息

	Description *string `json:"description,omitempty"`
	// 号码信息，如果为手机号，必须加上国家码。 例如中国大陆手机+86xxxxxxx，当填写手机号时 “country”参数必填,手机格式必须满足(^$|^[+]?[0-9]+$)

	SmsNumber *string `json:"smsNumber,omitempty"`
	// 若smsNumber为手机号,则需带上手机号所属的国家。 例如国家为中国大陆则country参数取值为chinaPR 国家和国家码的对应关系请参考：https://support.huaweicloud.com/api-meeting/meeting_21_0109.html

	Country *string `json:"country,omitempty"`
	// 邮箱地址

	EmailAddr *string `json:"emailAddr,omitempty"`
}

func (o VisionActiveCodeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VisionActiveCodeDto struct{}"
	}

	return strings.Join([]string{"VisionActiveCodeDto", string(data)}, " ")
}
