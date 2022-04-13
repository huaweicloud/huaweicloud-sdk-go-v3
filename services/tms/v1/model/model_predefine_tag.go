package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//  标签列表。
type PredefineTag struct {
	//   键。 最大长度36个字符。 字符集：A-Z，a-z ， 0-9，‘-’，‘_’，UNICODE字符（\\u4E00-\\u9FFF）。

	Key string `json:"key"`
	// 值。 每个值最大长度43个字符，可以为空字符串。 字符集：A-Z，a-z ， 0-9，‘.’，‘-’，‘_’，UNICODE字符（\\u4E00-\\u9FFF）。

	Value string `json:"value"`
	// 更新时间，采用UTC时间表示。2016-12-09T00:00:00Z

	UpdateTime *sdktime.SdkTime `json:"update_time"`
}

func (o PredefineTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PredefineTag struct{}"
	}

	return strings.Join([]string{"PredefineTag", string(data)}, " ")
}
