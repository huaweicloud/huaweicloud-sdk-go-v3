package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCceNodesLabelResponse Response Object
type ListCceNodesLabelResponse struct {

	// **参数解释**: 数据列表 **取值范围**: 取值1-100
	DataList *[]NodeLabelInfoResponse `json:"data_list,omitempty"`

	// **参数解释**: 集群NodeLabels总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListCceNodesLabelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCceNodesLabelResponse struct{}"
	}

	return strings.Join([]string{"ListCceNodesLabelResponse", string(data)}, " ")
}
