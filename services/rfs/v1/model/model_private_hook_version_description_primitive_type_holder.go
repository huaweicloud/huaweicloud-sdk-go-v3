package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateHookVersionDescriptionPrimitiveTypeHolder struct {

	// 私有hook版本的描述。可用于客户识别创建私有hook的版本。注意：hook版本为不可更新（immutable），所以该字段不可更新，如果需要更新，请删除后重建。
	HookVersionDescription *string `json:"hook_version_description,omitempty"`
}

func (o PrivateHookVersionDescriptionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateHookVersionDescriptionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateHookVersionDescriptionPrimitiveTypeHolder", string(data)}, " ")
}
