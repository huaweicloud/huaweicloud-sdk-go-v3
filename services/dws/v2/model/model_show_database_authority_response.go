package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDatabaseAuthorityResponse struct {

	// 对象权限集合
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
