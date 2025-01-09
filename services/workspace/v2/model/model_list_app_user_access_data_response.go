package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppUserAccessDataResponse Response Object
type ListAppUserAccessDataResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 用户统计指标
	Items          *[]AppUserAccessData `json:"items,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListAppUserAccessDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppUserAccessDataResponse struct{}"
	}

	return strings.Join([]string{"ListAppUserAccessDataResponse", string(data)}, " ")
}
