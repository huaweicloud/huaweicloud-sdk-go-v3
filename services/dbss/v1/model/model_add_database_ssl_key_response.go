package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDatabaseSslKeyResponse Response Object
type AddDatabaseSslKeyResponse struct {

	// 操作结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddDatabaseSslKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDatabaseSslKeyResponse struct{}"
	}

	return strings.Join([]string{"AddDatabaseSslKeyResponse", string(data)}, " ")
}
