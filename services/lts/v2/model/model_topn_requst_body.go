package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopnRequstBody struct {

	// 结束时间时间戳，毫秒时间
	EndTime int64 `json:"end_time"`

	// 是否降序  true / false
	IsDesc bool `json:"is_desc"`

	// 资源类型，log_group / log_stream / tenant
	ResourceType string `json:"resource_type"`

	// 排序依据，index/write/storage/basicTransfer/seniorTransfer/coldStorage，必须是search_list中存在的数据
	SortBy string `json:"sort_by"`

	// 开始时间时间戳，毫秒时间，最多支持30天范围内的查询
	StartTime int64 `json:"start_time"`

	// 查询前多少数据，范围1~100
	Topn int32 `json:"topn"`

	// 过滤条件 {     \"log_group_id\": \"xxxxxx\" }过滤器，为一个map结构，键为过滤属性，值为属性值，不支持模糊匹配
	Filter map[string]string `json:"filter"`

	// 查询数据类型，字符串数组可多种搭配，只能在index/write/storage/basicTransfer/seniorTransfer/coldStorage中选填
	SearchList []string `json:"search_list"`
}

func (o TopnRequstBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopnRequstBody struct{}"
	}

	return strings.Join([]string{"TopnRequstBody", string(data)}, " ")
}
