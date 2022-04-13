package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRomaAppResponse struct {
	// 总的数量

	Total *int32 `json:"total,omitempty"`
	// 当前页数量

	Size *int32 `json:"size,omitempty"`
	// 创建用户信息

	Apps           *[]ServerAppInfo `json:"apps,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListRomaAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRomaAppResponse struct{}"
	}

	return strings.Join([]string{"ListRomaAppResponse", string(data)}, " ")
}
