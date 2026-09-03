package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserInstanceListRequestBody 获取用户实例列表请求体
type ListUserInstanceListRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 页码
	PageNum *int32 `json:"page_num,omitempty"`

	// 查询记录数
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ListUserInstanceListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserInstanceListRequestBody struct{}"
	}

	return strings.Join([]string{"ListUserInstanceListRequestBody", string(data)}, " ")
}
