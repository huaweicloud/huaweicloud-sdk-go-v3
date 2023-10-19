package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAclAccountsRequest Request Object
type ListAclAccountsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ListAclAccountsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAclAccountsRequest struct{}"
	}

	return strings.Join([]string{"ListAclAccountsRequest", string(data)}, " ")
}
