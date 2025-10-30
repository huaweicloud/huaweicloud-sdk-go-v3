package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageBuildCommandRisksResponse Response Object
type ListImageBuildCommandRisksResponse struct {

	// **参数解释** 总数 **取值范围** 取值0-2147483547
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 风险列表 **取值范围**: 取值0-200
	DataList       *[]ImageBuildCommandRiskResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ListImageBuildCommandRisksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageBuildCommandRisksResponse struct{}"
	}

	return strings.Join([]string{"ListImageBuildCommandRisksResponse", string(data)}, " ")
}
