package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApiAclBindingV2Response Response Object
type DeleteApiAclBindingV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApiAclBindingV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApiAclBindingV2Response struct{}"
	}

	return strings.Join([]string{"DeleteApiAclBindingV2Response", string(data)}, " ")
}
