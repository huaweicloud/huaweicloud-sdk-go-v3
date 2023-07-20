package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOwnerRequestBody 修改数据库用户的请求body体。
type UpdateOwnerRequestBody struct {

	// 新owner名称。
	NewOwner string `json:"new_owner"`
}

func (o UpdateOwnerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOwnerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateOwnerRequestBody", string(data)}, " ")
}
