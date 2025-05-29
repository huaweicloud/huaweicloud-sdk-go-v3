package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskResultsDetailRequest Request Object
type ListTaskResultsDetailRequest struct {

	// 项目id
	ProjectUuid string `json:"project_uuid"`

	// 任务URI
	TaskUri string `json:"task_uri"`

	// 测试任务结果uri
	ResultUri string `json:"result_uri"`

	// 当前页数
	PageNo string `json:"page_no"`

	// 每页多少记录
	PageSize string `json:"page_size"`

	// 结果过滤条件
	Result *string `json:"result,omitempty"`
}

func (o ListTaskResultsDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskResultsDetailRequest struct{}"
	}

	return strings.Join([]string{"ListTaskResultsDetailRequest", string(data)}, " ")
}
