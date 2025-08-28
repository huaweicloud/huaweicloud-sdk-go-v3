package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TtscVoiceModelAssetMeta 音色模型元数据。
type TtscVoiceModelAssetMeta struct {

	// 音色性别。 * UNKNOW：中性音色 * MALE：男性音色 * FEMALE：女性音色  默认UNKNOW。
	Sex *string `json:"sex,omitempty"`

	ExternalVoiceMeta *TtscExternalVoiceAssetMeta `json:"external_voice_meta,omitempty"`

	// 音色语言。 * UNKNOW：未知 * CN：中文 * EN：英文 * GER：德语 * fr：法语 * Kr：韩语 * por：葡萄牙语 * JPN：日语 * Ita：意大利语 * ESP：西班牙语 * DBH：东北话 * GT：港台 * GXH：广西话 * HBH：湖北话 * SXH：陕西话 * SCH：四川话 * YY：粤语 * Russian: 俄罗斯语 * Filipino: 菲律宾语 * Dutch: 荷兰语 * Indonesian: 印尼语 * Vietnamese: 越南语 * Arabic: 阿拉伯语 * Turkish: 土耳其语 * Malay: 马来语 * Thai: 泰语 * Finnish: 芬兰语  默认UNKNOW。
	Language *string `json:"language,omitempty"`

	// 语速缩放比例
	SpeedRatio *float32 `json:"speed_ratio,omitempty"`

	// 音量缩放比例
	VolumeRatio *float32 `json:"volume_ratio,omitempty"`
}

func (o TtscVoiceModelAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TtscVoiceModelAssetMeta struct{}"
	}

	return strings.Join([]string{"TtscVoiceModelAssetMeta", string(data)}, " ")
}
