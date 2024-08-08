package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppConnectionResponse Response Object
type ListAppConnectionResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 应用使用记录列表。
	Items          *[]AppConnectionInfo `json:"items,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListAppConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppConnectionResponse struct{}"
	}

	return strings.Join([]string{"ListAppConnectionResponse", string(data)}, " ")
}
