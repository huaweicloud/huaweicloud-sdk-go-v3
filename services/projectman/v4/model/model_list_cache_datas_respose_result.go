package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCacheDatasResposeResult **参数解释：** 查询缓存的返回结果。
type ListCacheDatasResposeResult struct {

	// **参数解释：** 全部字段。
	Fields *[]FieldVo `json:"fields,omitempty"`

	// **参数解释：** 表头显示字段。
	VisibleFields *[]FieldVo `json:"visibleFields,omitempty"`
}

func (o ListCacheDatasResposeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCacheDatasResposeResult struct{}"
	}

	return strings.Join([]string{"ListCacheDatasResposeResult", string(data)}, " ")
}
