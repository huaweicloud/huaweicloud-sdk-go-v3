package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EngineModifyReq struct {

	// 变更的规格
	Flavor string `json:"flavor"`

	// 变更的配置，覆盖组件bp的input参数
	Inputs map[string]string `json:"inputs,omitempty"`
}

func (o EngineModifyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineModifyReq struct{}"
	}

	return strings.Join([]string{"EngineModifyReq", string(data)}, " ")
}
