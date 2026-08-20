package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceQuotaResult struct {

	// **参数解释**: 项目资源配额列表
	Resources *[]InstanceResourceQuotaResult `json:"resources,omitempty"`
}

func (o InstanceQuotaResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceQuotaResult struct{}"
	}

	return strings.Join([]string{"InstanceQuotaResult", string(data)}, " ")
}
