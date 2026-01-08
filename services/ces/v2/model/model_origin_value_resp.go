package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OriginValueResp **参数解释**： 实际维度信息。 **取值范围**： 字符串长度在 1 到 1024 之间。
type OriginValueResp struct {
}

func (o OriginValueResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OriginValueResp struct{}"
	}

	return strings.Join([]string{"OriginValueResp", string(data)}, " ")
}
