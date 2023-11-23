package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJdbcDriversResponse Response Object
type ListJdbcDriversResponse struct {

	// 驱动文件总数。
	Count *int32 `json:"count,omitempty"`

	// 驱动文件列表。
	Items          *[]DriverInfo `json:"items,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListJdbcDriversResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJdbcDriversResponse struct{}"
	}

	return strings.Join([]string{"ListJdbcDriversResponse", string(data)}, " ")
}
