package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceIdRequest Request Object
type ListInstanceIdRequest struct {

	// 域名ID
	DomainId string `json:"domain_id"`
}

func (o ListInstanceIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceIdRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceIdRequest", string(data)}, " ")
}
