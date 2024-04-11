package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportDesignModelTableDdlResultData 返回的数据信息。
type ExportDesignModelTableDdlResultData struct {

	// 接口导出的表的DDL语句。
	Value *string `json:"value,omitempty"`
}

func (o ExportDesignModelTableDdlResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDesignModelTableDdlResultData struct{}"
	}

	return strings.Join([]string{"ExportDesignModelTableDdlResultData", string(data)}, " ")
}
