package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParameterValuesInfo struct {

	// 参数名称。
	Name *string `json:"name,omitempty"`

	// 参数值。
	Value *string `json:"value,omitempty"`

	// 是否需要重启。 - false：否。 - true：是。
	RestartRequired *bool `json:"restart_required,omitempty"`

	// 是否只读。 - false：否。 - true：是。
	Readonly *bool `json:"readonly,omitempty"`

	// 参数值范围，如Integer取值0-1、Boolean取值true|false等。
	ValueRange *string `json:"value_range,omitempty"`

	// 参数类型,可取取值如下： - string - integer - boolean  - list  - float
	Type *string `json:"type,omitempty"`

	// 参数描述。
	Description *string `json:"description,omitempty"`
}

func (o ParameterValuesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParameterValuesInfo struct{}"
	}

	return strings.Join([]string{"ParameterValuesInfo", string(data)}, " ")
}
