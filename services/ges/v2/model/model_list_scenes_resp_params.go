package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListScenesRespParams struct {

	// 参数名称。
	Name *string `json:"name,omitempty"`

	// 参数类型。取值范围[\"string\",\"int\"]，目前仅支持\"string\"
	Type *string `json:"type,omitempty"`

	// 取值范围为空，或参数默认值，当为空是表示客户使用时必须传入此参数
	DefaultValue *string `json:"default_value,omitempty"`
}

func (o ListScenesRespParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScenesRespParams struct{}"
	}

	return strings.Join([]string{"ListScenesRespParams", string(data)}, " ")
}
