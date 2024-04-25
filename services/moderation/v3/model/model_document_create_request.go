package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DocumentCreateRequest struct {
	Data *DocumentCreateRequestData `json:"data"`

	// 事件类型，可选值如下： default：默认事件 liberal_arts_education：文科教育 sciences_education：理科教育 news：新闻 forums：论坛 novels：小说
	EventType string `json:"event_type"`

	// 文档中图片需要检测的风险类型，不传或为空时表示不审核图片内容， 可取值如下： politics: 涉政检测 porn：涉黄检测 terrorism: 暴恐检测 bad_scene: 不良场景检测 image_text: 图文检测
	ImageCategories *[]string `json:"image_categories,omitempty"`

	// 文档中文本需要检测的风险类型，不传或为空时表示不审核文本内容， 可取值如下： default: 检测涉政、暴恐、违禁、色情、辱骂、广告等违规内容
	TextCategories *[]string `json:"text_categories,omitempty"`

	// 网页视频中图片需要检测的风险类型，不传或为空时表示不审核网页视频内容， 可取值如下： politics: 涉政检测 porn：涉黄检测 terrorism: 暴恐检测 bad_scene: 不良场景检测 image_text: 图文检测
	VideoImageCategories *[]string `json:"video_image_categories,omitempty"`

	// 网页视频中音频需要检测的风险类型，不传或为空时表示不审核网页视频中音频的内容， 可取值如下： politics: 涉政检测 porn：涉黄检测 ad：广告检测 abuse：辱骂检测 moan：娇喘检测
	AudioCategories *[]string `json:"audio_categories,omitempty"`

	// 回调http接口：当该字段非空时，服务将根据该字段回调通知用户审核结果。
	Callback *string `json:"callback,omitempty"`

	// 用于回调通知时校验请求由华为云内容安全服务发起，由您自定义。随机字符串，由英文字母、数字、下划线组成，不超过64个字符。 说明：当seed非空时，headers中将包含X-Auth-Signature字段，字段的值使用HmacSHA256算法生成，待加密字符串由create_time、job_id、request_id、seed按照顺序拼接而成，密钥为seed。
	Seed *string `json:"seed,omitempty"`
}

func (o DocumentCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentCreateRequest struct{}"
	}

	return strings.Join([]string{"DocumentCreateRequest", string(data)}, " ")
}
