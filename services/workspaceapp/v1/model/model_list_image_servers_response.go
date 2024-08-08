package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageServersResponse Response Object
type ListImageServersResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 镜像实例列表返回列表条目数量上限为分页的最大上限值。
	Items          *[]ImageServer `json:"items,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListImageServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageServersResponse struct{}"
	}

	return strings.Join([]string{"ListImageServersResponse", string(data)}, " ")
}
