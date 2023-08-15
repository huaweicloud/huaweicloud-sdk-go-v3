package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityGroupResponse Response Object
type CreateSecurityGroupResponse struct {

	// 请求Id
	RequestId *string `json:"request_id,omitempty"`

	SecurityGroup  *SecurityGroupInfo `json:"security_group,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o CreateSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupResponse", string(data)}, " ")
}
