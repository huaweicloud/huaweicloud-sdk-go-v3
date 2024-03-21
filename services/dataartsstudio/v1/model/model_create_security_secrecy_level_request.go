package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecuritySecrecyLevelRequest Request Object
type CreateSecuritySecrecyLevelRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *CreateSecrecyLevelDto `json:"body,omitempty"`
}

func (o CreateSecuritySecrecyLevelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecuritySecrecyLevelRequest struct{}"
	}

	return strings.Join([]string{"CreateSecuritySecrecyLevelRequest", string(data)}, " ")
}
