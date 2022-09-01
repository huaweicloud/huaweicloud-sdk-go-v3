package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type TokenCatalog struct {

	// 该接口所属服务。
	Type string `json:"type" xml:"type"`

	// 服务ID。
	Id string `json:"id" xml:"id"`

	// 服务名称。
	Name string `json:"name" xml:"name"`

	// 终端节点。
	Endpoints []TokenCatalogEndpoint `json:"endpoints" xml:"endpoints"`
}

func (o TokenCatalog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TokenCatalog struct{}"
	}

	return strings.Join([]string{"TokenCatalog", string(data)}, " ")
}
