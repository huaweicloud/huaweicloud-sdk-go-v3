package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecuritySecrecyLevelRequest Request Object
type UpdateSecuritySecrecyLevelRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 密级id
	Id string `json:"id"`

	Body *UpdateSecrecyLevelDto `json:"body,omitempty"`
}

func (o UpdateSecuritySecrecyLevelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecuritySecrecyLevelRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecuritySecrecyLevelRequest", string(data)}, " ")
}
