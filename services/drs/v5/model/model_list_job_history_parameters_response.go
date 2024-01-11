package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobHistoryParametersResponse Response Object
type ListJobHistoryParametersResponse struct {

	// 历史记录总数。
	Count *int32 `json:"count,omitempty"`

	// 任务参数历史修改列表
	ParameterHistoryConfigList *[]ListJobHistoryParameter `json:"parameter_history_config_list,omitempty"`
	HttpStatusCode             int                        `json:"-"`
}

func (o ListJobHistoryParametersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobHistoryParametersResponse struct{}"
	}

	return strings.Join([]string{"ListJobHistoryParametersResponse", string(data)}, " ")
}
