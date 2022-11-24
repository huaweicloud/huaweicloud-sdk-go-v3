package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchAssociateDomainsV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAssociateDomainsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAssociateDomainsV2Response struct{}"
	}

	return strings.Join([]string{"BatchAssociateDomainsV2Response", string(data)}, " ")
}
