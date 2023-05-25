package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type MigrateDomainRequest struct {
	Body *DomainMigrate `json:"body,omitempty"`
}

func (o MigrateDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateDomainRequest struct{}"
	}

	return strings.Join([]string{"MigrateDomainRequest", string(data)}, " ")
}
