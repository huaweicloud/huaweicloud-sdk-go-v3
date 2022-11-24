package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddHooksResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *RepoHook `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddHooksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddHooksResponse struct{}"
	}

	return strings.Join([]string{"AddHooksResponse", string(data)}, " ")
}
