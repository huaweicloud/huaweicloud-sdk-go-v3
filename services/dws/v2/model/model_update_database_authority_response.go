package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDatabaseAuthorityResponse Response Object
type UpdateDatabaseAuthorityResponse struct {

	// **参数解释**： sql列表。 **取值范围**： 不涉及。
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
