package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApisBindedToAclPolicyV2Response Response Object
type ListApisBindedToAclPolicyV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次查询返回的API列表
	Apis           *[]AclBindApiInfo `json:"apis,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListApisBindedToAclPolicyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisBindedToAclPolicyV2Response struct{}"
	}

	return strings.Join([]string{"ListApisBindedToAclPolicyV2Response", string(data)}, " ")
}
