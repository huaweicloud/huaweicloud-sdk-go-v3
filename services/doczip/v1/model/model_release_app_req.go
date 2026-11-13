package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReleaseAppReq 应用发布请求消息
type ReleaseAppReq struct {

	// **参数解释**： 发布说明。 **约束限制**： 不涉及。 **取值范围**： [1-500]个字符。 **默认取值**： 不涉及。
	ReleaseNote string `json:"release_note"`
}

func (o ReleaseAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReleaseAppReq struct{}"
	}

	return strings.Join([]string{"ReleaseAppReq", string(data)}, " ")
}
