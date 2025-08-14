package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectConfigsResponse Response Object
type ListProjectConfigsResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 项目配置列表。
	Items          *[]ProjectConfig `json:"items,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListProjectConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectConfigsResponse", string(data)}, " ")
}
