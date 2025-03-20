package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncAddPersonnelRequest Request Object
type SyncAddPersonnelRequest struct {

	// 租户Id
	DomainId string `json:"domain_id"`
}

func (o SyncAddPersonnelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncAddPersonnelRequest struct{}"
	}

	return strings.Join([]string{"SyncAddPersonnelRequest", string(data)}, " ")
}
