package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterProtectPolicyTemplatesResponse Response Object
type ListClusterProtectPolicyTemplatesResponse struct {

	// **参数解释**: 策略模板总数 **取值范围**: 最小值0，最大值9223372036854775807
	TotalNum *int64 `json:"total_num,omitempty"`

	// **参数解释**: 策略模板列表 **取值范围**: 取值0-2147483647个PolicyTemplateInfo对象
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
