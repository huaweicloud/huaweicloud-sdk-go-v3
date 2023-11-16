package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBlockIpsRequest Request Object
type ListBlockIpsRequest struct {

	// 租户id
	DomainId string `json:"domain_id"`
}

func (o ListBlockIpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBlockIpsRequest struct{}"
	}

	return strings.Join([]string{"ListBlockIpsRequest", string(data)}, " ")
}
