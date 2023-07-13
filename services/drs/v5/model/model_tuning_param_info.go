package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TuningParamInfo 高级设置信息体。
type TuningParamInfo struct {

	// 全量调优参数。
	FullSync []TuningParameter `json:"full_sync"`

	// 增量抓取调优参数。
	IncreCapture []TuningParameter `json:"incre_capture"`

	// 增量回放调优参数。
	IncreApply []TuningParameter `json:"incre_apply"`

	// 增量日志拉取调优参数。
	IncreRelay []TuningParameter `json:"incre_relay"`
}

func (o TuningParamInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TuningParamInfo struct{}"
	}

	return strings.Join([]string{"TuningParamInfo", string(data)}, " ")
}
