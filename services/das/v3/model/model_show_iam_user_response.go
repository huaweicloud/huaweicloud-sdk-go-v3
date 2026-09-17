package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIamUserResponse Response Object
type ShowIamUserResponse struct {

	// IAM用户信息
	Data *interface{} `json:"data,omitempty"`

	// IAM用户的总数
	Total *int32 `json:"total,omitempty"`

	// 账号ID
	PrimaryAccountId *string `json:"primary_account_id,omitempty"`

	// 账号名称
	PrimaryAccountName *string `json:"primary_account_name,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o ShowIamUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIamUserResponse struct{}"
	}

	return strings.Join([]string{"ShowIamUserResponse", string(data)}, " ")
}
