package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateBlockchainCertByUserNameResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateBlockchainCertByUserNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBlockchainCertByUserNameResponse struct{}"
	}

	return strings.Join([]string{"CreateBlockchainCertByUserNameResponse", string(data)}, " ")
}
