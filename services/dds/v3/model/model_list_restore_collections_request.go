package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRestoreCollectionsRequest struct {
	// 语言。

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，可以调用“查询实例列表”接口获取。如果未申请实例，可以调用“创建实例”接口创建。

	InstanceId string `json:"instance_id"`
	// 数据库名称。

	DbName string `json:"db_name"`
	// 恢复时间点。UNIX时间戳格式，单位是毫秒，时区是UTC。

	RestoreTime string `json:"restore_time"`
	// 索引位置偏移量。取值大于或等于0。不传该参数时，查询偏移量默认为0。

	Offset *int32 `json:"offset,omitempty"`
	// 查询个数上限值。取值范围：1~100。不传该参数时，默认查询前100条信息。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRestoreCollectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreCollectionsRequest struct{}"
	}

	return strings.Join([]string{"ListRestoreCollectionsRequest", string(data)}, " ")
}
