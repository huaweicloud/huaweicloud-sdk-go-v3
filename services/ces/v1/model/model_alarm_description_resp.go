package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmDescriptionResp **参数解释**： 告警描述。 **取值范围**： 长度[0,256]个字符。
type AlarmDescriptionResp struct {
}

func (o AlarmDescriptionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmDescriptionResp struct{}"
	}

	return strings.Join([]string{"AlarmDescriptionResp", string(data)}, " ")
}
