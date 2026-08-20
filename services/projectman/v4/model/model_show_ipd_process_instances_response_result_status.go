package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpdProcessInstancesResponseResultStatus 评审单状态。
type ShowIpdProcessInstancesResponseResultStatus struct {

	// 状态编码。
	Code *string `json:"code,omitempty"`

	// 状态名称。
	Name *string `json:"name,omitempty"`
}

func (o ShowIpdProcessInstancesResponseResultStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpdProcessInstancesResponseResultStatus struct{}"
	}

	return strings.Join([]string{"ShowIpdProcessInstancesResponseResultStatus", string(data)}, " ")
}
