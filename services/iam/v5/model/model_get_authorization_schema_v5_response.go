package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAuthorizationSchemaV5Response Response Object
type GetAuthorizationSchemaV5Response struct {

	// 服务授权概要的版本号。
	Version *string `json:"version,omitempty"`

	// 云服务支持的授权项列表。
	Actions *[]Action `json:"actions,omitempty"`

	// 云服务支持的资源列表。
	Resources *[]Resource `json:"resources,omitempty"`

	// 云服务支持的条件键列表。
	Conditions *[]Condition `json:"conditions,omitempty"`

	// 云服务支持的操作列表。
	Operations     *[]Operation `json:"operations,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o GetAuthorizationSchemaV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAuthorizationSchemaV5Response struct{}"
	}

	return strings.Join([]string{"GetAuthorizationSchemaV5Response", string(data)}, " ")
}
