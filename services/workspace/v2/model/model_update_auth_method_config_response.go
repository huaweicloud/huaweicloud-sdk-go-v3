package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuthMethodConfigResponse Response Object
type UpdateAuthMethodConfigResponse struct {

	// 项目ID
	ProjectId      *string `json:"project_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAuthMethodConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthMethodConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuthMethodConfigResponse", string(data)}, " ")
}
