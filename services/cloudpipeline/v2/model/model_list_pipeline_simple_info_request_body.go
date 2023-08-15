package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelineSimpleInfoRequestBody 查询流水线信息对象
type ListPipelineSimpleInfoRequestBody struct {

	// 流水线名字。参数存在，则进行模糊匹配
	PipelineName *string `json:"pipeline_name,omitempty"`

	// devCloud项目ids。该参数存在，则获取对应项目下的流水线列表，逗号分隔，id个数取值[0,10]；如果不存在，则获取调用方所属租户的流水线列表
	ProjectIds *string `json:"project_ids,omitempty"`

	// 创建人id。该参数存在，逗号分隔，id个数取值[0,10]
	CreatorIds *string `json:"creator_ids,omitempty"`

	// 执行人id。该参数存在，逗号分隔，id个数取值[0,10]；
	ExecutorIds *string `json:"executor_ids,omitempty"`

	// 流水线运行状态。取值范围：waiting,running,verifying,handling,suspending,completed
	Status *string `json:"status,omitempty"`

	// 流水线结果，标记流水线。error、success、aborted
	Outcome *string `json:"outcome,omitempty"`

	// 用于排序的字段。取值为：pipeline_name,create_time,start_time
	SortKey *string `json:"sort_key,omitempty"`

	// 排序类型。asc按排序字段升序，desc按排序字段降序
	SortDir *string `json:"sort_dir,omitempty"`

	// codehub搜索链接：git@codehub.XXX.git；gitee搜索链接：git@gitee.com.XXX.git；github搜索链接：git@github.com.XXX.git等
	GitUrl *string `json:"git_url,omitempty"`

	// 偏移量。表示从此偏移量开始查询，offset大于等于0，默认取值为0
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的条目数量。取值[10-50]，默认取值为10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListPipelineSimpleInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineSimpleInfoRequestBody struct{}"
	}

	return strings.Join([]string{"ListPipelineSimpleInfoRequestBody", string(data)}, " ")
}
