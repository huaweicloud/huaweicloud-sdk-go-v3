package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowCurUserRoleRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id" xml:"project_id"`
}

func (o ShowCurUserRoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCurUserRoleRequest struct{}"
	}

	return strings.Join([]string{"ShowCurUserRoleRequest", string(data)}, " ")
}
