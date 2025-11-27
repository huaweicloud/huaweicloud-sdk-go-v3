package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateModuleVersionsResponse Response Object
type ListPrivateModuleVersionsResponse struct {

	// 私有模块版本的列表。默认以创建时间升序排序。
	Versions *[]PrivateModuleVersionSummary `json:"versions,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPrivateModuleVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateModuleVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateModuleVersionsResponse", string(data)}, " ")
}
