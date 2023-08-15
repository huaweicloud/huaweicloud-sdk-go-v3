package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemporaryAccessKeyByAgencyRequestBody
type CreateTemporaryAccessKeyByAgencyRequestBody struct {
	Auth *AgencyAuth `json:"auth"`
}

func (o CreateTemporaryAccessKeyByAgencyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemporaryAccessKeyByAgencyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTemporaryAccessKeyByAgencyRequestBody", string(data)}, " ")
}
