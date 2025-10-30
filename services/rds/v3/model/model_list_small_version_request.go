package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmallVersionRequest Request Object
type ListSmallVersionRequest struct {

	// 数据库引擎名。 取值范围： 支持的引擎如下，不区分大小写： PostgreSQL
	DatabaseName string `json:"database_name"`

	// 数据库版本号。（可输入小版本号）
	Version string `json:"version"`

	// 参数解释： 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。 约束限制：必须为数字，不能为负数。 取值范围：大于等于0的整数。 默认取值：0
	Offset *int32 `json:"offset,omitempty"`

	// 参数解释： 查询记录数。 约束限制：不涉及。 取值范围：默认为100，不能为负数，最小值为1，最大值为100。 默认取值：100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSmallVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmallVersionRequest struct{}"
	}

	return strings.Join([]string{"ListSmallVersionRequest", string(data)}, " ")
}
