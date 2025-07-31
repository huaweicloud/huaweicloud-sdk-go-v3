package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterProtectPolicyTemplatesResponse Response Object
type ListClusterProtectPolicyTemplatesResponse struct {

	// 策略模板总数
	TotalNum *int64 `json:"total_num,omitempty"`

	// 策略模板列表
	DataList       *[]PolicyTemplateInfo `json:"data_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListClusterProtectPolicyTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterProtectPolicyTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListClusterProtectPolicyTemplatesResponse", string(data)}, " ")
}
