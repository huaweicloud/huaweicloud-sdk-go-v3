package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportTopSqlListResponse Response Object
type ExportTopSqlListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportTopSqlListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTopSqlListResponse struct{}"
	}

	return strings.Join([]string{"ExportTopSqlListResponse", string(data)}, " ")
}
