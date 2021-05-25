package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddApplyJoinProjectForAgcRequest struct {
	// 租户id

	DomainId string `json:"Domain-Id"`
	// 用户id

	UserId string `json:"User-Id"`
	// 项目id

	ProjectId string `json:"project_id"`
}

func (o AddApplyJoinProjectForAgcRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddApplyJoinProjectForAgcRequest struct{}"
	}

	return strings.Join([]string{"AddApplyJoinProjectForAgcRequest", string(data)}, " ")
}
