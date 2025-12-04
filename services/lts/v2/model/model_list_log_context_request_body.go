package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogContextRequestBody 查询上下文日志请求体。
type ListLogContextRequestBody struct {

	// 日志单行序列号，字段值需要从日志结果中获取，纳秒级时间戳。
	LineNum *string `json:"line_num,omitempty"`

	// 自定义时间特性时间字段，字段值需要从日志结果中获取，毫秒级时间戳。若已开启云端结构化自定义时间功能，需要使用该字段和line_num字段共同进行上下文查询。
	Time *string `json:"__time__,omitempty"`

	// 指定起始日志往前（上文）的日志条数，取值范围 [0, 500] ，默认值100
	BackwardsSize *int32 `json:"backwards_size,omitempty"`

	// 指定起始日志往后（下文）的日志条数，取值范围 [0, 500] ，默认值100
	ForwardsSize *int32 `json:"forwards_size,omitempty"`

	// 分页查询时，若上次查询的返回结果中包含该字段，需要增加该字段参与分页查询。
	ScrollId *string `json:"scroll_id,omitempty"`
}

func (o ListLogContextRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogContextRequestBody struct{}"
	}

	return strings.Join([]string{"ListLogContextRequestBody", string(data)}, " ")
}
