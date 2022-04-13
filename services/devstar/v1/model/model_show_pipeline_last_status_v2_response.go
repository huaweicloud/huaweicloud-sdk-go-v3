package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPipelineLastStatusV2Response struct {
	// 流水线id

	PipelineId *string `json:"pipeline_id,omitempty"`
	// 流水线名称

	Name *string `json:"name,omitempty"`
	// 执行状态

	Status *string `json:"status,omitempty"`
	// 执行结果

	Result *string `json:"result,omitempty"`
	// 执行人

	Executor *string `json:"executor,omitempty"`
	// 启动时间

	StartTime *string `json:"start_time,omitempty"`
	// 结束时间

	FinishTime *string `json:"finish_time,omitempty"`
	// 运行详情链接

	DetailUrl *string `json:"detail_url,omitempty"`
	// 编辑链接

	ModifyUrl *string `json:"modify_url,omitempty"`
	// 流水线执行序号

	BuildId *string `json:"build_id,omitempty"`
	// 阶段信息

	Stages         *[]PipelineStageResp `json:"stages,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowPipelineLastStatusV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineLastStatusV2Response struct{}"
	}

	return strings.Join([]string{"ShowPipelineLastStatusV2Response", string(data)}, " ")
}
