package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserJdbcDriversResponse Response Object
type ListUserJdbcDriversResponse struct {

	// 驱动文件总数。
	Count *int32 `json:"count,omitempty"`

	// 驱动文件列表。
	Items          *[]DriverInfo `json:"items,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListUserJdbcDriversResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserJdbcDriversResponse struct{}"
	}

	return strings.Join([]string{"ListUserJdbcDriversResponse", string(data)}, " ")
}
