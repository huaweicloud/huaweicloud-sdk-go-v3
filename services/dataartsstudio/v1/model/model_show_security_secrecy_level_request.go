package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecuritySecrecyLevelRequest Request Object
type ShowSecuritySecrecyLevelRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 密级id
	Id string `json:"id"`
}

func (o ShowSecuritySecrecyLevelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecuritySecrecyLevelRequest struct{}"
	}

	return strings.Join([]string{"ShowSecuritySecrecyLevelRequest", string(data)}, " ")
}
