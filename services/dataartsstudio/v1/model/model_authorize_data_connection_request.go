package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeDataConnectionRequest Request Object
type AuthorizeDataConnectionRequest struct {

	// 需要授权的数据连接id。
	DataConnectionId string `json:"data_connection_id"`

	// 需要授权的工作空间id。
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o AuthorizeDataConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeDataConnectionRequest struct{}"
	}

	return strings.Join([]string{"AuthorizeDataConnectionRequest", string(data)}, " ")
}
