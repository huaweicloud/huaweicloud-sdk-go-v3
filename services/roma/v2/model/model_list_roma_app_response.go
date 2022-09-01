package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRomaAppResponse struct {

	// 总的数量
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 当前页数量
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 创建用户信息
	Apps           *[]ServerAppInfo `json:"apps,omitempty" xml:"apps"`
	HttpStatusCode int              `json:"-"`
}

func (o ListRomaAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRomaAppResponse struct{}"
	}

	return strings.Join([]string{"ListRomaAppResponse", string(data)}, " ")
}
