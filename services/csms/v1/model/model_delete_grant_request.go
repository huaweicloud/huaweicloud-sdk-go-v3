package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGrantRequest Request Object
type DeleteGrantRequest struct {

	// 被授权用户ID
	GranteeUser *string `json:"grantee_user,omitempty"`

	// 被授权资源ID
	ResourceId string `json:"resource_id"`
}

func (o DeleteGrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGrantRequest struct{}"
	}

	return strings.Join([]string{"DeleteGrantRequest", string(data)}, " ")
}
