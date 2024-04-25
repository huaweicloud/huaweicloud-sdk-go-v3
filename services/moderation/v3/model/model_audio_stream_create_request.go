package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AudioStreamCreateRequest struct {
	Data *AudioStreamCreateRequestData `json:"data"`

	// 事件类型，可选值如下： default：默认事件 audiobook：有声书 education：教育音频 game：游戏语音房 live：秀场直播 ecommerce：电商直播 voiceroom：交友语音房 private：私密语音聊天
	EventType string `json:"event_type"`

	// 需要检测的风险类型，列表不能为空。 porn：涉黄检测 politics: 涉政检测 abuse: 辱骂检测 ad: 广告检测 moan: 娇喘检测
	Categories []string `json:"categories"`

	// 回调http接口，服务将根据该字段回调通知用户审核结果，流未结束时，回调审核违规内容，流结束时，审核片段违规或不违规都将回调客户端。 回调内容如下： ``` {     \"job_id\":\"xxxxxx\",     \"status\":\"running\", //running - 审核中，succeeded - 审核完成，failed - 审核失败     \"request_id\":\"2419446b1fe14203f64e4018d12db3dd\",     \"create_time\":\"2022-07-30T08:57:11.011Z\",     \"update_time\":\"2022-07-30T08:57:14.014Z\",     \"result\":{         \"suggestion\":\"block\",         \"details\":[             {                 \"suggestion\":\"block\",                 \"label\":\"politics\",                 \"audio_text\":\"xxxx\",                 \"start_time\":\"2022-07-30T08:57:11.011Z\", // 当前音频片段开始的绝对时间                 \"end_time\":\"2022-07-30T08:57:21.011Z\",     // 当前音频片段结束的绝对时间                 \"segments\":[                     {                         \"segment\":\"xxx\"                     },                     {                         \"segment\":\"xxx\"                     },                     {                         \"segment\":\"xxx\"                     }                 ]             }         ],         \"request_params\":{             \"event_type\":\"default\",             \"data\":{                 \"url\":\"rtmp://xxxx\"             },             \"callback\":\"http://xxx\",             \"categories\":[                 \"porn\",                 \"ad\"             ]         }     } }
	Callback string `json:"callback"`

	// 用于回调通知时校验请求由华为云内容安全服务发起，由您自定义。随机字符串，由英文字母、数字、下划线组成，不超过64个字符。 说明：当seed非空时，headers中将包含X-Auth-Signature字段，字段的值使用HmacSHA256算法生成，待加密字符串由create_time、job_id、request_id、seed按照顺序拼接而成，密钥为seed。
	Seed *string `json:"seed,omitempty"`
}

func (o AudioStreamCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioStreamCreateRequest struct{}"
	}

	return strings.Join([]string{"AudioStreamCreateRequest", string(data)}, " ")
}
