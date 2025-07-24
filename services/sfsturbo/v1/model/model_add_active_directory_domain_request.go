package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddActiveDirectoryDomainRequest Request Object
type AddActiveDirectoryDomainRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	Body *AddActiveDirectoryDomainRequestBody `json:"body,omitempty"`
}

func (o AddActiveDirectoryDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddActiveDirectoryDomainRequest struct{}"
	}

	return strings.Join([]string{"AddActiveDirectoryDomainRequest", string(data)}, " ")
}
