package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2EndpointName **参数解释**：端点名称。 **取值范围**：支持小写字母、数字、中划线，必须字母开头、字母或数字结尾，长度为1~63个字符。
type V2EndpointName struct {
}

func (o V2EndpointName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2EndpointName struct{}"
	}

	return strings.Join([]string{"V2EndpointName", string(data)}, " ")
}
