package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostStatusResponse Response Object
type ListHostStatusResponse struct {

	// **参数解释**: 总数 **取值范围**: 取值0-2097152
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 查询弹性云服务器状态列表 **取值范围**: 不涉及
	DataList       *[]Host `json:"data_list,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListHostStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostStatusResponse struct{}"
	}

	return strings.Join([]string{"ListHostStatusResponse", string(data)}, " ")
}
