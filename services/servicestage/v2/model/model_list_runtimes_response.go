package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRuntimesResponse Response Object
type ListRuntimesResponse struct {

	// 运行时列表。
	Runtimes       *[]RuntimeTypeView `json:"runtimes,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListRuntimesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuntimesResponse struct{}"
	}

	return strings.Join([]string{"ListRuntimesResponse", string(data)}, " ")
}
