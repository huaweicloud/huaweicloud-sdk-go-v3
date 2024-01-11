package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobParametersRequest Request Object
type ListJobParametersRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset 大于等于 0。默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。默认为10，取值范围【1-1000】
	Limit *int32 `json:"limit,omitempty"`

	// 根据参数名查询。
	Name *string `json:"name,omitempty"`
}

func (o ListJobParametersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobParametersRequest struct{}"
	}

	return strings.Join([]string{"ListJobParametersRequest", string(data)}, " ")
}
