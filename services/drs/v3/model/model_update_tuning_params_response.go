package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateTuningParamsResponse struct {

	// 全量调优参数
	FullSync *[]TuningParameter `json:"full_sync,omitempty" xml:"full_sync"`

	// 增量抓取调优参数
	IncreCapture *[]TuningParameter `json:"incre_capture,omitempty" xml:"incre_capture"`

	// 增量回放调优参数
	IncreApply *[]TuningParameter `json:"incre_apply,omitempty" xml:"incre_apply"`

	// 增量日志拉取调优参数
	IncreRelay *[]TuningParameter `json:"incre_relay,omitempty" xml:"incre_relay"`

	// 参数修改是否成功
	ModifyResult   *string `json:"modify_result,omitempty" xml:"modify_result"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTuningParamsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTuningParamsResponse struct{}"
	}

	return strings.Join([]string{"UpdateTuningParamsResponse", string(data)}, " ")
}
