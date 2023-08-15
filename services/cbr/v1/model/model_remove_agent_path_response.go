package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveAgentPathResponse Response Object
type RemoveAgentPathResponse struct {

	// 移除的路径列表
	Removed *[]string `json:"removed,omitempty"`

	// 不存在的路径
	NotExisted     *[]string `json:"not_existed,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o RemoveAgentPathResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveAgentPathResponse struct{}"
	}

	return strings.Join([]string{"RemoveAgentPathResponse", string(data)}, " ")
}
