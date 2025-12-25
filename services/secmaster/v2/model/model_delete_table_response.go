package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTableResponse Response Object
type DeleteTableResponse struct {

	// 表ID
	TableId *string `json:"table_id,omitempty"`

	ProcessStatus *IsapTableProcessStatus `json:"process_status,omitempty"`

	// 毫秒时间戳
	DeleteTime     *int64 `json:"delete_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTableResponse struct{}"
	}

	return strings.Join([]string{"DeleteTableResponse", string(data)}, " ")
}
