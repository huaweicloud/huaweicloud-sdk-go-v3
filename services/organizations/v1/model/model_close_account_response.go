package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloseAccountResponse Response Object
type CloseAccountResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CloseAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseAccountResponse struct{}"
	}

	return strings.Join([]string{"CloseAccountResponse", string(data)}, " ")
}
