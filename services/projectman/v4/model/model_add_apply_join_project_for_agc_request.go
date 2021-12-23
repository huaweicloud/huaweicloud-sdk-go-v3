package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddApplyJoinProjectForAgcRequest struct {
	// 租户id

	DomainId string `json:"Domain-Id"`
	// 用户id

	UserId string `json:"User-Id"`
	// devcloud的项目id

	ProjectId string `json:"project_id"`
}

func (o AddApplyJoinProjectForAgcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddApplyJoinProjectForAgcRequest struct{}"
	}

	return strings.Join([]string{"AddApplyJoinProjectForAgcRequest", string(data)}, " ")
}
