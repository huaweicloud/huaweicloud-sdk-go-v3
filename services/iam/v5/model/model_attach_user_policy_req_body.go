package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachUserPolicyReqBody Contains information about a id of a user.
type AttachUserPolicyReqBody struct {

	// IAM用户ID。
	UserId string `json:"user_id"`
}

func (o AttachUserPolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachUserPolicyReqBody struct{}"
	}

	return strings.Join([]string{"AttachUserPolicyReqBody", string(data)}, " ")
}
