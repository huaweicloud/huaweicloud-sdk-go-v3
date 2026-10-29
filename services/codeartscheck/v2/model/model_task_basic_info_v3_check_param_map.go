package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskBasicInfoV3CheckParamMap **参数解释：** 自定义参数列表checkParamList的map格式，比如含有自定义参数key为param1，值为value1，则checkParamMap含有字段param1，值为value1。
type TaskBasicInfoV3CheckParamMap struct {

	// **参数解释：** 字段的name为自定义参数的Key，字段的值为自定义参数的Value。 **取值范围：** 长度为[1~8192]个字符。
	Param1 *string `json:"param1,omitempty"`

	// **参数解释：** 字段的name为自定义参数的Key，字段的值为自定义参数的Value。 **取值范围：** 长度为[1~8192]个字符。
	Param2 *string `json:"param2,omitempty"`
}

func (o TaskBasicInfoV3CheckParamMap) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskBasicInfoV3CheckParamMap struct{}"
	}

	return strings.Join([]string{"TaskBasicInfoV3CheckParamMap", string(data)}, " ")
}
