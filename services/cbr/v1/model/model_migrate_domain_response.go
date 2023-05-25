package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type MigrateDomainResponse struct {
	Body           map[string]string `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o MigrateDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateDomainResponse struct{}"
	}

	return strings.Join([]string{"MigrateDomainResponse", string(data)}, " ")
}
