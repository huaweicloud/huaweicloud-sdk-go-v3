package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowActiveDirectoryDomainRequest Request Object
type ShowActiveDirectoryDomainRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`
}

func (o ShowActiveDirectoryDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowActiveDirectoryDomainRequest struct{}"
	}

	return strings.Join([]string{"ShowActiveDirectoryDomainRequest", string(data)}, " ")
}
