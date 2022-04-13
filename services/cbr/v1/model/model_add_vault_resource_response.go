package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddVaultResourceResponse struct {
	// 已添加的资源ID列表

	AddResourceIds *[]string `json:"add_resource_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AddVaultResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVaultResourceResponse struct{}"
	}

	return strings.Join([]string{"AddVaultResourceResponse", string(data)}, " ")
}
