package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServersResponse Response Object
type ListServersResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 服务器列表返回列表条目数量上限为分页的最大上限值。
	Items          *[]AppServer `json:"items,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersResponse struct{}"
	}

	return strings.Join([]string{"ListServersResponse", string(data)}, " ")
}
