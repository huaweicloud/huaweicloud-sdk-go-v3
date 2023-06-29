package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecretResponse Response Object
type CreateSecretResponse struct {

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	Secret         *CreateUpdateSecretRespSecret `json:"secret,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o CreateSecretResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretResponse struct{}"
	}

	return strings.Join([]string{"CreateSecretResponse", string(data)}, " ")
}
