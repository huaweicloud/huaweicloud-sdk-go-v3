package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunLanguageDetectionResponse Response Object
type RunLanguageDetectionResponse struct {

	// 调用成功时表示调用结果，语种缩写对应名称如下： 阿拉伯语 ar 爱沙尼亚语 et 保加利亚语 bg 冰岛语 is 波兰语 pl 波斯尼亚语 bs 波斯语 fa 丹麦语 da 德语 de 俄语 ru 法语 fr 芬兰语 fi 高棉语 km 韩语 ko 加泰罗尼亚语 ca 捷克语 cs 克罗地亚语 hr 拉脱维亚语 lv 立陶宛语 lt 罗马尼亚语 ro 马耳他语 mt 马来西亚语 ms 马其顿语 mk 孟加拉语 bn 缅甸语 my 南非荷兰语 af 挪威语 no 葡萄牙语 pt 日语 ja 瑞典语 sv 塞尔维亚语 sr 斯洛伐克语 sk 斯洛文尼亚语 sl 斯瓦希里语 sw 泰语 th 土耳其语 tr 威尔士语 cy 乌尔都语 ur 乌克兰语 uk 西班牙语 es 希伯来语 he 希腊语 el 匈牙利语 hu 意大利语 it 印地语 hi 印尼语 id 英语 en 越南语 vi 中文 zh 无法识别语种 unk 当输入文本过短或不明确时，识别结果可能不准确； 当输入文本包含多种语言时，会返回占比最高的语种。 调用失败时无此字段。
	DetectedLanguage *string `json:"detected_language,omitempty"`

	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。
	ErrorCode *string `json:"error_code,omitempty"`

	// 调用失败时的错误信息。调用成功时无此字段。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunLanguageDetectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunLanguageDetectionResponse struct{}"
	}

	return strings.Join([]string{"RunLanguageDetectionResponse", string(data)}, " ")
}
