package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkspaceResponse Response Object
type CreateWorkspaceResponse struct {

	// 返回的数据信息。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateWorkspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkspaceResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkspaceResponse", string(data)}, " ")
}
