package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebFrameworkHostInfoResponse Response Object
type ListWebFrameworkHostInfoResponse struct {

	// **参数解释** 总数 **取值范围** 最小值0，最大值10000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 服务器列表 **取值范围** 最小值0，最大值10000
	DataList       *[]WebFrameworkHostInfo `json:"data_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListWebFrameworkHostInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebFrameworkHostInfoResponse struct{}"
	}

	return strings.Join([]string{"ListWebFrameworkHostInfoResponse", string(data)}, " ")
}
