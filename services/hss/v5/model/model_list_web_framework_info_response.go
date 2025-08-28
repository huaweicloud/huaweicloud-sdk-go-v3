package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebFrameworkInfoResponse Response Object
type ListWebFrameworkInfoResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: Web框架列表 **取值范围**: 最小值0，最大值10000
	DataList       *[]WebFrameworkInfo `json:"data_list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListWebFrameworkInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebFrameworkInfoResponse struct{}"
	}

	return strings.Join([]string{"ListWebFrameworkInfoResponse", string(data)}, " ")
}
