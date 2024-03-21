package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecuritySecrecyLevelRequest Request Object
type DeleteSecuritySecrecyLevelRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 密级id
	Id string `json:"id"`
}

func (o DeleteSecuritySecrecyLevelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecuritySecrecyLevelRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecuritySecrecyLevelRequest", string(data)}, " ")
}
