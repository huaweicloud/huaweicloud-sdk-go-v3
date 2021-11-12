package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDatabaseUsersRequest struct {
	// 实例ID，可以调用“查询实例列表”接口获取。如果未申请实例，可以调用“创建实例”接口创建。

	InstanceId string `json:"instance_id"`
	// 用户名称。 - 长度为1~64位，可以包含大写字母（A~Z）、小写字母（a~z）、数字（0~9）、中划线、下划线和点。

	UserName *string `json:"user_name,omitempty"`
	// 数据库名称，默认为admin。 - 长度为1~64位，可以包含大写字母（A~Z）、小写字母（a~z）、数字（0~9）、下划线。

	DbName *string `json:"db_name,omitempty"`
	// 索引位置偏移量。 取值大于或等于0。不传该参数时，查询偏移量默认为0。

	Offset *int32 `json:"offset,omitempty"`
	// 查询实例个数上限值。 取值范围：1~100。不传该参数时，默认查询前100条实例信息。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDatabaseUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseUsersRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseUsersRequest", string(data)}, " ")
}
