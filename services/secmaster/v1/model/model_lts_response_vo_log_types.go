package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LtsResponseVoLogTypes 日志类型列表
type LtsResponseVoLogTypes struct {

	// 类型前缀
	SecmLts *[]string `json:"secm_lts_,omitempty"`
}

func (o LtsResponseVoLogTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsResponseVoLogTypes struct{}"
	}

	return strings.Join([]string{"LtsResponseVoLogTypes", string(data)}, " ")
}
