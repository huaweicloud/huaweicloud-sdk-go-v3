package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiMoveParaDto struct {

	// 父目录编号
	TargetPid *string `json:"target_pid,omitempty"`

	// 需要移动的目录
	Apis *[]string `json:"apis,omitempty"`
}

func (o ApiMoveParaDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiMoveParaDto struct{}"
	}

	return strings.Join([]string{"ApiMoveParaDto", string(data)}, " ")
}
