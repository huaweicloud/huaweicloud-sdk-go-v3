package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubmodulesResponse Response Object
type ListSubmodulesResponse struct {

	// 子模块列表
	Body *[]SubmoduleDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSubmodulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubmodulesResponse struct{}"
	}

	return strings.Join([]string{"ListSubmodulesResponse", string(data)}, " ")
}
