package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRecordResponse Response Object
type CheckRecordResponse struct {

	// 历史记录数据
	Data *[]CheckRecordDataInfo `json:"data,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CheckRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRecordResponse struct{}"
	}

	return strings.Join([]string{"CheckRecordResponse", string(data)}, " ")
}
