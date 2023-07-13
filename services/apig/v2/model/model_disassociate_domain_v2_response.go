package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateDomainV2Response Response Object
type DisassociateDomainV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateDomainV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateDomainV2Response struct{}"
	}

	return strings.Join([]string{"DisassociateDomainV2Response", string(data)}, " ")
}
