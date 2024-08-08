package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateHookDescriptionPrimitiveTypeHolder struct {

	// 私有hook的描述。可用于客户识别创建的私有hook。可通过UpdatePrivateHook API更新私有hook的描述。
	HookDescription *string `json:"hook_description,omitempty"`
}

func (o PrivateHookDescriptionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateHookDescriptionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateHookDescriptionPrimitiveTypeHolder", string(data)}, " ")
}
