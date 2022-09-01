package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddApplyJoinProjectForAgcRequest struct {

	// 租户id
	DomainId string `json:"Domain-Id" xml:"Domain-Id"`

	// 用户id
	UserId string `json:"User-Id" xml:"User-Id"`

	// devcloud项目的32位id
	ProjectId string `json:"project_id" xml:"project_id"`
}

func (o AddApplyJoinProjectForAgcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddApplyJoinProjectForAgcRequest struct{}"
	}

	return strings.Join([]string{"AddApplyJoinProjectForAgcRequest", string(data)}, " ")
}
