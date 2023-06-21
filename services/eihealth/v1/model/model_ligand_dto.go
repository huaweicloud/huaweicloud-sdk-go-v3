package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配体配置
type LigandDto struct {
	Ligand *DrugFile `json:"ligand"`

	// 计算个数
	Count int32 `json:"count"`
}

func (o LigandDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LigandDto struct{}"
	}

	return strings.Join([]string{"LigandDto", string(data)}, " ")
}
