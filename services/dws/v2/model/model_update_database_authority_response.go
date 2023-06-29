package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDatabaseAuthorityResponse Response Object
type UpdateDatabaseAuthorityResponse struct {

	// sql列表
	ViewSql        *[]string `json:"view_sql,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateDatabaseAuthorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseAuthorityResponse struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseAuthorityResponse", string(data)}, " ")
}
