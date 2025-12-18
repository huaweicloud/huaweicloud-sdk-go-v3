package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFullSqlListResponse Response Object
type ExportFullSqlListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportFullSqlListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFullSqlListResponse struct{}"
	}

	return strings.Join([]string{"ExportFullSqlListResponse", string(data)}, " ")
}
