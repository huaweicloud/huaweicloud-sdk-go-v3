package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTableSchemaRequest struct {

	// 表ID。
	TableId string `json:"table_id" xml:"table_id"`
}

func (o ShowTableSchemaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableSchemaRequest struct{}"
	}

	return strings.Join([]string{"ShowTableSchemaRequest", string(data)}, " ")
}
