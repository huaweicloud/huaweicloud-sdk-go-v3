package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListStackOutputsResponse struct {

	// 堆栈输出
	Outputs *[]StackOutput `json:"outputs,omitempty"`

	// 下一页的标记信息
	NextMarker     *string `json:"next_marker,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListStackOutputsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStackOutputsResponse struct{}"
	}

	return strings.Join([]string{"ListStackOutputsResponse", string(data)}, " ")
}
