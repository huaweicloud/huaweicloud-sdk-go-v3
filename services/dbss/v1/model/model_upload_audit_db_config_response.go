package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAuditDbConfigResponse Response Object
type UploadAuditDbConfigResponse struct {

	// 操作结果  - success：成功  - failed：失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadAuditDbConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAuditDbConfigResponse struct{}"
	}

	return strings.Join([]string{"UploadAuditDbConfigResponse", string(data)}, " ")
}
