package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyHotfixesRequestBody struct {

	// **参数解释**: 热补丁属性信息。 **约束限制**: 不涉及。
	Hotfixes *interface{} `json:"hotfixes"`
}

func (o ModifyHotfixesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyHotfixesRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyHotfixesRequestBody", string(data)}, " ")
}
