package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListLoginRecordsNewResponse struct {

	// 用户登录记录总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 用户登录记录。
	Records        *[]Record `json:"records,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListLoginRecordsNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoginRecordsNewResponse struct{}"
	}

	return strings.Join([]string{"ListLoginRecordsNewResponse", string(data)}, " ")
}
