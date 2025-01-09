package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyDetailInfoByIdRequest Request Object
type ListPolicyDetailInfoByIdRequest struct {

	// 策略组id。
	PolicyGroupId string `json:"policy_group_id"`
}

func (o ListPolicyDetailInfoByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyDetailInfoByIdRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyDetailInfoByIdRequest", string(data)}, " ")
}
