package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
