package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnscopedTokenInfoCatalog struct {

	// 终端节点ID。
	Id *string `json:"id,omitempty"`

	// 服务名称。
	Name *string `json:"name,omitempty"`

	// 该接口所属服务。
	Type *string `json:"type,omitempty"`

	Endpoints *[]UnscopedTokenInfoEndpoints `json:"endpoints,omitempty"`
}

func (o UnscopedTokenInfoCatalog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnscopedTokenInfoCatalog struct{}"
	}

	return strings.Join([]string{"UnscopedTokenInfoCatalog", string(data)}, " ")
}
