package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAgreementResponse Response Object
type CreateAgreementResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateAgreementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgreementResponse struct{}"
	}

	return strings.Join([]string{"CreateAgreementResponse", string(data)}, " ")
}
