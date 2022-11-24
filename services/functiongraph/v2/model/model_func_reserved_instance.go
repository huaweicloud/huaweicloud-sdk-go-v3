package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FuncReservedInstance struct {

	// 函数urn
	FuncUrn string `json:"func_urn"`

	// 预留实例数目
	Count int64 `json:"count"`
}

func (o FuncReservedInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FuncReservedInstance struct{}"
	}

	return strings.Join([]string{"FuncReservedInstance", string(data)}, " ")
}
