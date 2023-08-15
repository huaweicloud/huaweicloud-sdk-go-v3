package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkspacesResponse Response Object
type DeleteWorkspacesResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteWorkspacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkspacesResponse struct{}"
	}

	return strings.Join([]string{"DeleteWorkspacesResponse", string(data)}, " ")
}
