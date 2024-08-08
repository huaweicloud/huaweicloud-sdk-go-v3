package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateHookDefaultVersionPrimitiveTypeHolder struct {

	// 私有hook的默认版本号。版本号必须遵循语义化版本号（Semantic Version），为用户自定义。
	DefaultVersion *string `json:"default_version,omitempty"`
}

func (o PrivateHookDefaultVersionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateHookDefaultVersionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateHookDefaultVersionPrimitiveTypeHolder", string(data)}, " ")
}
