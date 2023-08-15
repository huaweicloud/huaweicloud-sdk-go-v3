package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessPolicyObjectsResponse Response Object
type ListAccessPolicyObjectsResponse struct {

	// 查询接入策略应用对象响应。
	PolicyObjectsList *[]AccessPolicyObject `json:"policy_objects_list,omitempty"`

	// 对象总数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAccessPolicyObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPolicyObjectsResponse struct{}"
	}

	return strings.Join([]string{"ListAccessPolicyObjectsResponse", string(data)}, " ")
}
