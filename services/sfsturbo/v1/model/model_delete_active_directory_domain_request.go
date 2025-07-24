package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteActiveDirectoryDomainRequest Request Object
type DeleteActiveDirectoryDomainRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	Body *DeleteActiveDirectoryDomainRequestBody `json:"body,omitempty"`
}

func (o DeleteActiveDirectoryDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteActiveDirectoryDomainRequest struct{}"
	}

	return strings.Join([]string{"DeleteActiveDirectoryDomainRequest", string(data)}, " ")
}
