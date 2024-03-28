package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PreviewTableRequest Request Object
type PreviewTableRequest struct {

	// 待预览的表所在的数据库名称。
	DatabaseName string `json:"database_name"`

	// 待预览的表名称。
	TableName string `json:"table_name"`

	// 预览表的模式，取值为““SYNC””或者““ASYNC””默认值为：“SYNC”。
	Mode *string `json:"mode,omitempty"`
}

func (o PreviewTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviewTableRequest struct{}"
	}

	return strings.Join([]string{"PreviewTableRequest", string(data)}, " ")
}
