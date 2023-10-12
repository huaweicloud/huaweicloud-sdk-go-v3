package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSwitchoverTestRecordsResponse Response Object
type ListSwitchoverTestRecordsResponse struct {

	// 操作请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 倒换测试记录信息列表
	SwitchoverTestRecords *[]SwitchoverTestRecord `json:"switchover_test_records,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSwitchoverTestRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSwitchoverTestRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListSwitchoverTestRecordsResponse", string(data)}, " ")
}
