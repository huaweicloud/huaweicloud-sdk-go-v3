package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServiceSpecificCredentialV5Response Response Object
type DeleteServiceSpecificCredentialV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServiceSpecificCredentialV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceSpecificCredentialV5Response struct{}"
	}

	return strings.Join([]string{"DeleteServiceSpecificCredentialV5Response", string(data)}, " ")
}
