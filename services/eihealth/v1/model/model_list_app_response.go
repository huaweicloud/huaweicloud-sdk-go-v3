package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAppResponse struct {

	// 应用列表
	Apps *[]AppListDto `json:"apps,omitempty"`

	// 应用总条数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppResponse struct{}"
	}

	return strings.Join([]string{"ListAppResponse", string(data)}, " ")
}
