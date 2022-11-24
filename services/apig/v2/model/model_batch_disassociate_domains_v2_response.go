package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDisassociateDomainsV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDisassociateDomainsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisassociateDomainsV2Response struct{}"
	}

	return strings.Join([]string{"BatchDisassociateDomainsV2Response", string(data)}, " ")
}
