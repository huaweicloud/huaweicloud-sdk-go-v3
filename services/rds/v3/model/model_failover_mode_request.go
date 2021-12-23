package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FailoverModeRequest struct {
	// 数据库主备同步模式

	Mode string `json:"mode"`
}

func (o FailoverModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailoverModeRequest struct{}"
	}

	return strings.Join([]string{"FailoverModeRequest", string(data)}, " ")
}
