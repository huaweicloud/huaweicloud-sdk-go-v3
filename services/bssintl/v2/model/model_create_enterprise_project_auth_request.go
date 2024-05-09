package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnterpriseProjectAuthRequest Request Object
type CreateEnterpriseProjectAuthRequest struct {
}

func (o CreateEnterpriseProjectAuthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnterpriseProjectAuthRequest struct{}"
	}

	return strings.Join([]string{"CreateEnterpriseProjectAuthRequest", string(data)}, " ")
}
