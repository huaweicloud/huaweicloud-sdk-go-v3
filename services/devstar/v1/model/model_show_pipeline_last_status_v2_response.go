package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPipelineLastStatusV2Response struct {

	// 流水线id
	PipelineId *string `json:"pipeline_id,omitempty" xml:"pipeline_id"`

	// 流水线名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 执行状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 执行结果
	Result *string `json:"result,omitempty" xml:"result"`

	// 执行人
	Executor *string `json:"executor,omitempty" xml:"executor"`

	// 启动时间
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 结束时间
	FinishTime *string `json:"finish_time,omitempty" xml:"finish_time"`

	// 运行详情链接
	DetailUrl *string `json:"detail_url,omitempty" xml:"detail_url"`

	// 编辑链接
	ModifyUrl *string `json:"modify_url,omitempty" xml:"modify_url"`

	// 流水线执行序号
	BuildId *string `json:"build_id,omitempty" xml:"build_id"`

	// 阶段信息
	Stages         *[]PipelineStageResp `json:"stages,omitempty" xml:"stages"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowPipelineLastStatusV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineLastStatusV2Response struct{}"
	}

	return strings.Join([]string{"ShowPipelineLastStatusV2Response", string(data)}, " ")
}
