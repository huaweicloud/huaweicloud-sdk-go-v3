package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserConnectionResponse Response Object
type ListUserConnectionResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 应用使用记录列表。
	Items          *[]UserConnectionInfo `json:"items,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListUserConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserConnectionResponse struct{}"
	}

	return strings.Join([]string{"ListUserConnectionResponse", string(data)}, " ")
}
