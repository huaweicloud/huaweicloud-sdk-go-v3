package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateObsTargetPolicyResponse Response Object
type UpdateObsTargetPolicyResponse struct {

	// 绑定关系ID
	TargetId *string `json:"target_id,omitempty"`

	Policy *ObsDataRepositoryPolicy `json:"policy,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateObsTargetPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateObsTargetPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateObsTargetPolicyResponse", string(data)}, " ")
}
