package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTableContentRequest Request Object
type ShowTableContentRequest struct {

	// 待预览的表所在的数据库名称。
	DatabaseName string `json:"database_name"`

	// 待预览的表名称。
	TableName string `json:"table_name"`

	// 预览表的模式，取值为““SYNC””或者““ASYNC””默认值为：“SYNC”。
	Mode *string `json:"mode,omitempty"`
}

func (o ShowTableContentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableContentRequest struct{}"
	}

	return strings.Join([]string{"ShowTableContentRequest", string(data)}, " ")
}
