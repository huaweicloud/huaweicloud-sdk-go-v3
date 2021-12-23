package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowCurUserRoleRequest struct {
	// devcloud的项目id

	ProjectId string `json:"project_id"`
}

func (o ShowCurUserRoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCurUserRoleRequest struct{}"
	}

	return strings.Join([]string{"ShowCurUserRoleRequest", string(data)}, " ")
}
