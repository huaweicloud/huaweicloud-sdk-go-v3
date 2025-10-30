package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTuningParamsResponse Response Object
type UpdateTuningParamsResponse struct {

	// 全量调优参数
	FullSync *[]TuningParameter `json:"full_sync,omitempty"`

	// 增量抓取调优参数
	IncreCapture *[]TuningParameter `json:"incre_capture,omitempty"`

	// 增量回放调优参数
	IncreApply *[]TuningParameter `json:"incre_apply,omitempty"`

	// 增量日志拉取调优参数
	IncreRelay *[]TuningParameter `json:"incre_relay,omitempty"`

	// 参数修改是否成功，查询参数时不返回该字段。
	ModifyResult   *string `json:"modify_result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTuningParamsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTuningParamsResponse struct{}"
	}

	return strings.Join([]string{"UpdateTuningParamsResponse", string(data)}, " ")
}
