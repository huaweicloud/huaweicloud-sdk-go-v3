package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHooksResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *RepoListHook `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListHooksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHooksResponse struct{}"
	}

	return strings.Join([]string{"ListHooksResponse", string(data)}, " ")
}
