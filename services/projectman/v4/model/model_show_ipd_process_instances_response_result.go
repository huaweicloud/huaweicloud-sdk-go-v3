package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpdProcessInstancesResponseResult 请求结果。
type ShowIpdProcessInstancesResponseResult struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 评审单列表。
	ProcessInstances *[]ShowIpdProcessInstancesResponseResultProcessInstances `json:"process_instances,omitempty"`
}

func (o ShowIpdProcessInstancesResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpdProcessInstancesResponseResult struct{}"
	}

	return strings.Join([]string{"ShowIpdProcessInstancesResponseResult", string(data)}, " ")
}
