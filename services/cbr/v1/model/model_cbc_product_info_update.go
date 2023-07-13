package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CbcProductInfoUpdate struct {

	// 产品标识，通过订购询价接口获得，长度限制：1-64，只能由字母、数字、“_”、“-”组成。。
	ProductId string `json:"productId"`

	// 资源容量大小，取值范围：10-10485760
	ResourceSize int32 `json:"resourceSize"`

	// 资源容量度量标识，枚举值17：GB
	ResourceSizeMeasureId *int32 `json:"resourceSizeMeasureId,omitempty"`

	// 用户购买云服务产品的资源规格 Enum: [vault.backup.server.normal，vault.backup.turbo.normal, vault.backup.database.normal，vault.backup.volume.normal，vault.backup.rds.normal，vault.replication.server.normal，vault.hybrid.server.normal]
	ResourceSpecCode string `json:"resourceSpecCode"`
}

func (o CbcProductInfoUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbcProductInfoUpdate struct{}"
	}

	return strings.Join([]string{"CbcProductInfoUpdate", string(data)}, " ")
}
