package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VarOverridesPrimitiveTypeHolder struct {
	VarOverrides *VarOverridesPrimitiveTypeHolderVarOverrides `json:"var_overrides,omitempty"`
}

func (o VarOverridesPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VarOverridesPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"VarOverridesPrimitiveTypeHolder", string(data)}, " ")
}
