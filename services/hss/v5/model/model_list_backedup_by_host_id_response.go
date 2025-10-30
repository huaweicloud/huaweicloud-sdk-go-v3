package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackedupByHostIdResponse Response Object
type ListBackedupByHostIdResponse struct {

	// **参数解释** 总数 **取值范围** 取值0-2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 备份列表 **取值范围** 取值0-200
	DataList       *[]BackCopyByHostIdResponse `json:"data_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListBackedupByHostIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackedupByHostIdResponse struct{}"
	}

	return strings.Join([]string{"ListBackedupByHostIdResponse", string(data)}, " ")
}
