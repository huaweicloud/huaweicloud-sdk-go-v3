package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkspaceResponse Response Object
type UpdateWorkspaceResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateWorkspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkspaceResponse struct{}"
	}

	return strings.Join([]string{"UpdateWorkspaceResponse", string(data)}, " ")
}
