package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSlowSqlListResponse Response Object
type ExportSlowSqlListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportSlowSqlListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSlowSqlListResponse struct{}"
	}

	return strings.Join([]string{"ExportSlowSqlListResponse", string(data)}, " ")
}
