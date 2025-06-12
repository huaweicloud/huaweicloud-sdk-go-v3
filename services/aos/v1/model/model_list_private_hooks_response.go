package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateHooksResponse Response Object
type ListPrivateHooksResponse struct {
	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 私有hook的列表。默认以创建时间降序排序。
	Hooks          *[]PrivateHookSummary `json:"hooks,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListPrivateHooksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateHooksResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateHooksResponse", string(data)}, " ")
}
