package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateHookVersionsResponse Response Object
type ListPrivateHookVersionsResponse struct {

	// 私有hook version的列表。默认以创建时间降序排序。
	Versions *[]PrivateHookVersionSummary `json:"versions,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPrivateHookVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateHookVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateHookVersionsResponse", string(data)}, " ")
}
