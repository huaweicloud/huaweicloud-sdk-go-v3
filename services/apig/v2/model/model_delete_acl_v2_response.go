package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAclV2Response Response Object
type DeleteAclV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAclV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAclV2Response struct{}"
	}

	return strings.Join([]string{"DeleteAclV2Response", string(data)}, " ")
}
