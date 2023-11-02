package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MoleculeFileDto struct {
	File *MoleculeFile `json:"file"`

	// 分子个数
	Count int32 `json:"count"`
}

func (o MoleculeFileDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoleculeFileDto struct{}"
	}

	return strings.Join([]string{"MoleculeFileDto", string(data)}, " ")
}
