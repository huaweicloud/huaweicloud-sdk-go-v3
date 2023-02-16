package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAuthorizationRequest struct {

	// 事件id
	CaseId string `json:"case_id"`

	Body *CreateUserCenterAuthorizationV2Req `json:"body,omitempty"`
}

func (o CreateAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"CreateAuthorizationRequest", string(data)}, " ")
}
