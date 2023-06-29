package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnterpriseProjectAuthResponse Response Object
type CreateEnterpriseProjectAuthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateEnterpriseProjectAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnterpriseProjectAuthResponse struct{}"
	}

	return strings.Join([]string{"CreateEnterpriseProjectAuthResponse", string(data)}, " ")
}
