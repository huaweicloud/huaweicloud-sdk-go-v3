package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApisUnbindedToAclPolicyV2Response Response Object
type ListApisUnbindedToAclPolicyV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次查询返回的API列表
	Apis           *[]UnbindApiForAcl `json:"apis,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListApisUnbindedToAclPolicyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisUnbindedToAclPolicyV2Response struct{}"
	}

	return strings.Join([]string{"ListApisUnbindedToAclPolicyV2Response", string(data)}, " ")
}
