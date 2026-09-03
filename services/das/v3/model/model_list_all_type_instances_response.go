package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllTypeInstancesResponse Response Object
type ListAllTypeInstancesResponse struct {

	// 错误信息条件列表
	EntryNames     *[]string `json:"entry_names,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAllTypeInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllTypeInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListAllTypeInstancesResponse", string(data)}, " ")
}
