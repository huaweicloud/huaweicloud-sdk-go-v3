package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSlowSqlResponse Response Object
type ExportSlowSqlResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportSlowSqlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSlowSqlResponse struct{}"
	}

	return strings.Join([]string{"ExportSlowSqlResponse", string(data)}, " ")
}
