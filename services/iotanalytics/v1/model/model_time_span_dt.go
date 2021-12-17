package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 起止时间
type TimeSpanDt struct {
	// 起始时间, 使用UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss.SSS'Z',示例：2021-02-01T00:00:00.123Z

	From *string `json:"from,omitempty"`
	// 结束时间，使用UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss.SSS'Z',示例：2021-02-01T00:00:00.123Z

	To *string `json:"to,omitempty"`
}

func (o TimeSpanDt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeSpanDt struct{}"
	}

	return strings.Join([]string{"TimeSpanDt", string(data)}, " ")
}
