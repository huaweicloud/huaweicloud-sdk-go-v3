package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAuditDbConfigRequest Request Object
type UploadAuditDbConfigRequest struct {
}

func (o UploadAuditDbConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAuditDbConfigRequest struct{}"
	}

	return strings.Join([]string{"UploadAuditDbConfigRequest", string(data)}, " ")
}
