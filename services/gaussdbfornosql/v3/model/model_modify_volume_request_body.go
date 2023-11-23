package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyVolumeRequestBody struct {

	// 待变更到的磁盘容量。单位GB，取值为整数。 扩容场景下，必须大于当前磁盘容量。 缩容场景下，必须大于已用量的125%，向上取整。 磁盘容量的上下限与所选引擎类型以及规格相关。   - [GeminiDB Cassandra请参见[数据库实例规格](https://support.huaweicloud.com/cassandraug-nosql/nosql_05_0017.html)。](tag:hc)   - [GeminiDB Cassandra请参见[数据库实例规格。](https://support.huaweicloud.com/intl/zh-cn/cassandraug-nosql/nosql_05_0017.html)](tag:hk)   - [GeminiDB Redis请参见[数据库实例规格。](https://support.huaweicloud.com/redisug-nosql/nosql_05_0059.html)](tag:hc)   - [GeminiDB Redis请参见[数据库实例规格。](https://support.huaweicloud.com/intl/zh-cn/redisug-nosql/nosql_05_0059.html)](tag:hk)
	Size int32 `json:"size"`

	// 扩容包年包月实例存储容量时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。   - true，表示自动从账户中支付。   - false，表示手动从账户中支付，默认为该方式。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o ModifyVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyVolumeRequestBody", string(data)}, " ")
}
