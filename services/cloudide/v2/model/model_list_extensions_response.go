package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExtensionsResponse Response Object
type ListExtensionsResponse struct {

	// 插件列表查询结果集合
	Results        *[]ExtensionQueryResult `json:"results,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListExtensionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExtensionsResponse struct{}"
	}

	return strings.Join([]string{"ListExtensionsResponse", string(data)}, " ")
}
