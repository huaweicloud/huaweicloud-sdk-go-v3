package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityLevelFromEntityRequest Request Object
type DeleteSecurityLevelFromEntityRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 资产id
	Guid string `json:"guid"`
}

func (o DeleteSecurityLevelFromEntityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityLevelFromEntityRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityLevelFromEntityRequest", string(data)}, " ")
}
