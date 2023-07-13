package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VisionActiveCodeDto struct {

	// 终端类型。 - idea-hub：智能协作大屏 - huawei-vision：智慧屏TV - welink-desktop(iwb)：SmartRooms会议版 - smart-rooms：SmartRooms完整版
	DevType string `json:"devType"`

	// 部门编码，若不携带则默认根部门。
	DeptCode *string `json:"deptCode,omitempty"`

	// 终端的名称。
	DevName string `json:"devName"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// 号码信息，如果为手机号，必须加上国家码。 例如中国大陆手机+86xxxxxxx，当填写手机号时 “country”参数必填,手机格式必须满足(^$|^[+]?[0-9]+$)
	SmsNumber *string `json:"smsNumber,omitempty"`

	// [[手机号所属的国家](https://support.huaweicloud.com/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hws)[[手机号所属的国家](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hk) 。
	Country *string `json:"country,omitempty"`

	// 邮箱地址。 > 如无中国大陆手机号，则邮箱必填。
	EmailAddr *string `json:"emailAddr,omitempty"`
}

func (o VisionActiveCodeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VisionActiveCodeDto struct{}"
	}

	return strings.Join([]string{"VisionActiveCodeDto", string(data)}, " ")
}
