package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyTuningParamsReq 修改调优参数请求体
type ModifyTuningParamsReq struct {

	// 全量参数的参数名及对应的取值
	FullSync map[string]string `json:"full_sync,omitempty"`

	// 增量抓取参数的参数名及对应的取值
	IncreCapture map[string]string `json:"incre_capture,omitempty"`

	// 增量回放参数的参数名及对应的取值
	IncreApply map[string]string `json:"incre_apply,omitempty"`

	// 增量日志拉取参数的参数名及对应的取值 -slotAdvanceInterval： 源库逻辑复制槽推进间隔
	IncreRelay map[string]string `json:"incre_relay,omitempty"`

	// 初始化参数，首次调用时需要传true，其他时候不传。
	Recovery *bool `json:"recovery,omitempty"`
}

func (o ModifyTuningParamsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTuningParamsReq struct{}"
	}

	return strings.Join([]string{"ModifyTuningParamsReq", string(data)}, " ")
}
