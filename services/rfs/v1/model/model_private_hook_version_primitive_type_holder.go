package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateHookVersionPrimitiveTypeHolder struct {

	// 私有hook的版本号。版本号必须遵循语义化版本号（Semantic Version），为用户自定义。
	HookVersion string `json:"hook_version"`
}

func (o PrivateHookVersionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateHookVersionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateHookVersionPrimitiveTypeHolder", string(data)}, " ")
}
