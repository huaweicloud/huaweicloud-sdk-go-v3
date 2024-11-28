package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRedisPitrRestoreTimeResponse Response Object
type ListRedisPitrRestoreTimeResponse struct {

	// Redis可恢复时间点列表。 yyyy-mm-ddThh:mm:ssZ字符串格式，T指某个时间的开始，Z指时区偏移量。
	RestoreTime *[]string `json:"restore_time,omitempty"`

	// Redis实例可恢复时间点总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRedisPitrRestoreTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRedisPitrRestoreTimeResponse struct{}"
	}

	return strings.Join([]string{"ListRedisPitrRestoreTimeResponse", string(data)}, " ")
}
