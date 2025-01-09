package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScreenRecordsResponse Response Object
type ListScreenRecordsResponse struct {

	// 录屏记录。
	ScreenRecords *[]QueryScreenRecordDetailRsp `json:"screen_records,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListScreenRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScreenRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListScreenRecordsResponse", string(data)}, " ")
}
