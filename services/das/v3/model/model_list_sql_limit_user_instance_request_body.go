package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlLimitUserInstanceRequestBody 获取用户实例列表请求体
type ListSqlLimitUserInstanceRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 页码，取值范围 [0, 5000]，默认为 0
	PageNum *int32 `json:"page_num,omitempty"`

	// 查询记录数，取值范围 [0, 100]，默认为 20
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ListSqlLimitUserInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlLimitUserInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"ListSqlLimitUserInstanceRequestBody", string(data)}, " ")
}
