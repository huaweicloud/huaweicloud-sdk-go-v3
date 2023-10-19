package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAclAccountRemarkRequestBody 更新ACL账号备注请求
type UpdateAclAccountRemarkRequestBody struct {

	// 备注信息
	Description *string `json:"description,omitempty"`
}

func (o UpdateAclAccountRemarkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclAccountRemarkRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAclAccountRemarkRequestBody", string(data)}, " ")
}
