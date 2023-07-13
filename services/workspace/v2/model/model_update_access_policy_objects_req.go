package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAccessPolicyObjectsReq struct {

	// 更新策略应用对象列表请求。
	PolicyObjectsList *[]AccessPolicyObjectInfo `json:"policy_objects_list,omitempty"`
}

func (o UpdateAccessPolicyObjectsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessPolicyObjectsReq struct{}"
	}

	return strings.Join([]string{"UpdateAccessPolicyObjectsReq", string(data)}, " ")
}
