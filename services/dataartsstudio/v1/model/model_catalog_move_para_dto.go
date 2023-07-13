package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CatalogMoveParaDto struct {

	// 父目录编号
	TargetPid *string `json:"target_pid,omitempty"`
}

func (o CatalogMoveParaDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogMoveParaDto struct{}"
	}

	return strings.Join([]string{"CatalogMoveParaDto", string(data)}, " ")
}
