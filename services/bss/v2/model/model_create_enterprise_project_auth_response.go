package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
