package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckEnterpriseIdResponse Response Object
type CheckEnterpriseIdResponse struct {

	// 企业ID是否已被使用。
	IsUsed *bool `json:"is_used,omitempty"`

	// 企业ID。
	EnterpriseId   *string `json:"enterprise_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckEnterpriseIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckEnterpriseIdResponse struct{}"
	}

	return strings.Join([]string{"CheckEnterpriseIdResponse", string(data)}, " ")
}
