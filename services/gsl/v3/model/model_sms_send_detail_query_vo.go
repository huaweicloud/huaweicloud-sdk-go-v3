package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmsSendDetailQueryVo struct {

	// 容器ID
	Cid *string `json:"cid,omitempty"`

	// MSISDN
	Msisdn *string `json:"msisdn,omitempty"`

	// 发送时间
	SentTime *sdktime.SdkTime `json:"sent_time,omitempty"`

	// 接收时间
	ReceivedTime *sdktime.SdkTime `json:"received_time,omitempty"`

	// 短信状态:1发送中 2.已送达 3.失败
	SmsStatus *int32 `json:"sms_status,omitempty"`

	// 短信内容
	SmsContent *string `json:"sms_content,omitempty"`
}

func (o SmsSendDetailQueryVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsSendDetailQueryVo struct{}"
	}

	return strings.Join([]string{"SmsSendDetailQueryVo", string(data)}, " ")
}
