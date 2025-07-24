package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateActiveDirectoryDomainRequest Request Object
type UpdateActiveDirectoryDomainRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	Body *UpdateActiveDirectoryDomainRequestBody `json:"body,omitempty"`
}

func (o UpdateActiveDirectoryDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateActiveDirectoryDomainRequest struct{}"
	}

	return strings.Join([]string{"UpdateActiveDirectoryDomainRequest", string(data)}, " ")
}
