package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobParametersResponse Response Object
type ListJobParametersResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 任务参数列表
	ParameterConfigList *[]ParameterConfig `json:"parameter_config_list,omitempty"`
	HttpStatusCode      int                `json:"-"`
}

func (o ListJobParametersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobParametersResponse struct{}"
	}

	return strings.Join([]string{"ListJobParametersResponse", string(data)}, " ")
}
