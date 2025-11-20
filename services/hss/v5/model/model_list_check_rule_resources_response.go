package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCheckRuleResourcesResponse Response Object
type ListCheckRuleResourcesResponse struct {

	// 受配置检测影响的镜像数据总量
	TotalNum *int32 `json:"total_num,omitempty"`

	// 数据列表
	DataList       *[]CheckRuleImageResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListCheckRuleResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCheckRuleResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListCheckRuleResourcesResponse", string(data)}, " ")
}
