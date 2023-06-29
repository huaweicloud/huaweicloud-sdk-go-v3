package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecretResponse Response Object
type UpdateSecretResponse struct {

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	Secret         *CreateUpdateSecretRespSecret `json:"secret,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o UpdateSecretResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecretResponse", string(data)}, " ")
}
