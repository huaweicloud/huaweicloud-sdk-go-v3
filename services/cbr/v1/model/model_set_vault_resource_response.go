package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SetVaultResourceResponse struct {

	// 本次设置的资源id列表。
	SetResourceIds *[]string `json:"set_resource_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o SetVaultResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetVaultResourceResponse struct{}"
	}

	return strings.Join([]string{"SetVaultResourceResponse", string(data)}, " ")
}
