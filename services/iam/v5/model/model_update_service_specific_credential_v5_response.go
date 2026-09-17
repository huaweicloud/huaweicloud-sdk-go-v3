package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServiceSpecificCredentialV5Response Response Object
type UpdateServiceSpecificCredentialV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateServiceSpecificCredentialV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceSpecificCredentialV5Response struct{}"
	}

	return strings.Join([]string{"UpdateServiceSpecificCredentialV5Response", string(data)}, " ")
}
