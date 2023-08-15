package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyTuningParams 修改高级设置信息体。
type ModifyTuningParams struct {

	// 全量参数的参数名及对应的取值。
	FullSync map[string]string `json:"full_sync,omitempty"`

	// 增量抓取参数的参数名及对应的取值。
	IncreCapture map[string]string `json:"incre_capture"`

	// 增量回放参数的参数名及对应的取值。
	IncreApply map[string]string `json:"incre_apply"`

	// 增量日志拉取参数的参数名及对应的取值。
	IncreRelay map[string]string `json:"incre_relay,omitempty"`

	// 是否一键还原到默认值。
	Recovery *bool `json:"recovery,omitempty"`
}

func (o ModifyTuningParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTuningParams struct{}"
	}

	return strings.Join([]string{"ModifyTuningParams", string(data)}, " ")
}
