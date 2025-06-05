package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKeystorePermissionRequest Request Object
type ShowKeystorePermissionRequest struct {

	// 文件秘钥Id
	KeystoreId string `json:"keystore_id"`

	// **参数解释**： 每页显示的条目数量。 **约束限制**： 不涉及。 **取值范围**： 只能使用数字，小于等于100。
	PageSize int32 `json:"page_size"`

	// 分页页码，表示从此页开始查询，page大于等于1
	Page string `json:"page"`
}

func (o ShowKeystorePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeystorePermissionRequest struct{}"
	}

	return strings.Join([]string{"ShowKeystorePermissionRequest", string(data)}, " ")
}
