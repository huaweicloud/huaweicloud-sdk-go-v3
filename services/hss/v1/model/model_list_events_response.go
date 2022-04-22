package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEventsResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 查询弹性云服务器状态列表
	DataList       *[]Event `json:"data_list,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsResponse struct{}"
	}

	return strings.Join([]string{"ListEventsResponse", string(data)}, " ")
}
