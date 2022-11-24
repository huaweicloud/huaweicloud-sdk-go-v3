package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProjectModulesResponse struct {

	// 模块总数
	Total *int32 `json:"total,omitempty"`

	// 模块列表
	Modules        *[]ProjectModule `json:"modules,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListProjectModulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectModulesResponse struct{}"
	}

	return strings.Join([]string{"ListProjectModulesResponse", string(data)}, " ")
}
