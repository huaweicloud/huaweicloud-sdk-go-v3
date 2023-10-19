package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAclAccountRequest Request Object
type CreateAclAccountRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *CreateAclAccountRequestBody `json:"body,omitempty"`
}

func (o CreateAclAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAclAccountRequest struct{}"
	}

	return strings.Join([]string{"CreateAclAccountRequest", string(data)}, " ")
}
