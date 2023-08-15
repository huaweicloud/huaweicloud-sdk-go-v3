package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateSecurityLevelToEntitieRequest Request Object
type AssociateSecurityLevelToEntitieRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 资产id
	Guid string `json:"guid"`

	// 资产密级
	SecurityLevel string `json:"security-level"`
}

func (o AssociateSecurityLevelToEntitieRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSecurityLevelToEntitieRequest struct{}"
	}

	return strings.Join([]string{"AssociateSecurityLevelToEntitieRequest", string(data)}, " ")
}
