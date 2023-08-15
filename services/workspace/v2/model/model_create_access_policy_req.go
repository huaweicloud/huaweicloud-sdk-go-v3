package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAccessPolicyReq struct {
	Policy *AccessPolicyInfo `json:"policy"`

	// 策略应用对象列表。
	PolicyObjectsList *[]AccessPolicyObjectInfo `json:"policy_objects_list,omitempty"`
}

func (o CreateAccessPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessPolicyReq struct{}"
	}

	return strings.Join([]string{"CreateAccessPolicyReq", string(data)}, " ")
}
