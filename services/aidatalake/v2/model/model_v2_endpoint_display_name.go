package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2EndpointDisplayName **参数解释**：端点显示名称。 **取值范围**：长度为1~63个字符，支持字母、数字、中划线、汉字。
type V2EndpointDisplayName struct {
}

func (o V2EndpointDisplayName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2EndpointDisplayName struct{}"
	}

	return strings.Join([]string{"V2EndpointDisplayName", string(data)}, " ")
}
