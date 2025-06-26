package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDomainNameRequest Request Object
type AddDomainNameRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	Body *AddDomainNameRequestBody `json:"body,omitempty"`
}

func (o AddDomainNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDomainNameRequest struct{}"
	}

	return strings.Join([]string{"AddDomainNameRequest", string(data)}, " ")
}
