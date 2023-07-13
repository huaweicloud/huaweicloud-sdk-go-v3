package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSqlConversionResponse Response Object
type RunSqlConversionResponse struct {

	// 是否支持SQL语句转换。
	IsSupportConversion *bool `json:"is_support_conversion,omitempty"`

	// 转换后的SQL语句。
	ConvertedSqlStatement *string `json:"converted_sql_statement,omitempty"`

	// 不支持SQL语句转换的详情。
	UnsupportedItems *[]UnSupportedItem `json:"unsupported_items,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o RunSqlConversionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSqlConversionResponse struct{}"
	}

	return strings.Join([]string{"RunSqlConversionResponse", string(data)}, " ")
}
