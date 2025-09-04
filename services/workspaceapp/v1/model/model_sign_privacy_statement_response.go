package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SignPrivacyStatementResponse Response Object
type SignPrivacyStatementResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SignPrivacyStatementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignPrivacyStatementResponse struct{}"
	}

	return strings.Join([]string{"SignPrivacyStatementResponse", string(data)}, " ")
}
