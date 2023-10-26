package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckMigrationConnectivityResponse Response Object
type CheckMigrationConnectivityResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckMigrationConnectivityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckMigrationConnectivityResponse struct{}"
	}

	return strings.Join([]string{"CheckMigrationConnectivityResponse", string(data)}, " ")
}
