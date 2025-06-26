package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerNameResponse Response Object
type UpdateServerNameResponse struct {

	// 修改指定ID的源端服务器信息成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateServerNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateServerNameResponse", string(data)}, " ")
}
