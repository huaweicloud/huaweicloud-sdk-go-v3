package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatabaseAuthorityResponse Response Object
type ShowDatabaseAuthorityResponse struct {

	// **参数解释**： 对象权限集合。 **取值范围**： 不涉及。
	Authorities    *[]ObjectAuthority `json:"authorities,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowDatabaseAuthorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseAuthorityResponse struct{}"
	}

	return strings.Join([]string{"ShowDatabaseAuthorityResponse", string(data)}, " ")
}
