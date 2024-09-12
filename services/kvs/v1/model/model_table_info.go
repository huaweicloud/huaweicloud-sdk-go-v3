package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableInfo struct {

	// 表状态。 - \"creating\" - \"active\" - \"deleting\"
	TableStatus *string `bson:"table_status,omitempty"`
}

func (o TableInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableInfo struct{}"
	}

	return strings.Join([]string{"TableInfo", string(data)}, " ")
}
