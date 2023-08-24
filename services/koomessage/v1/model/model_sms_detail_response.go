package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmsDetailResponse struct {

	// 短信接收方的号码。
	To string `json:"to"`

	// 短信发送方的号码。
	ChannelNum string `json:"channel_num"`

	// 短信的唯一标识。
	MsgId string `json:"msg_id"`

	// 短信状态码。  以下举例状态码及其说明，具体处理建议请参考[API错误码](https://support.huaweicloud.com/api-msgsms/sms_05_0050.html)。  - 000000：短信平台处理请求成功 - E200015：待发送短信数量太大 - E200028：模板变量校验失败 - E200029：模板类型校验失败 - E200030：模板未激活 - E200031：协议校验失败 - E200033：模板类型不正确 - E200041：同一短信内容接收号码重复
	SendStatus string `json:"send_status"`

	// 短信资源的创建时间。  即短信平台接收到用户发送短信请求的时间，为UTC时间。 格式为：yyyy-MM-dd'T'HH:mm:ss'Z'。
	CreateTime string `json:"create_time"`
}

func (o SmsDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsDetailResponse struct{}"
	}

	return strings.Join([]string{"SmsDetailResponse", string(data)}, " ")
}
