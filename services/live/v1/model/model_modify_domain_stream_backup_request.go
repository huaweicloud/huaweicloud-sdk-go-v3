package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyDomainStreamBackupRequest Request Object
type ModifyDomainStreamBackupRequest struct {
	Body *DomainStreamBackupInfo `json:"body,omitempty"`
}

func (o ModifyDomainStreamBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDomainStreamBackupRequest struct{}"
	}

	return strings.Join([]string{"ModifyDomainStreamBackupRequest", string(data)}, " ")
}
