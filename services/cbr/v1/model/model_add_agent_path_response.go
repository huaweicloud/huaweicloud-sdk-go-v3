package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddAgentPathResponse struct {

	// 新添加成功的路径列表
	Added *[]string `json:"added,omitempty"`

	// 已经存在的路径列表
	Existed        *[]string `json:"existed,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AddAgentPathResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAgentPathResponse struct{}"
	}

	return strings.Join([]string{"AddAgentPathResponse", string(data)}, " ")
}
