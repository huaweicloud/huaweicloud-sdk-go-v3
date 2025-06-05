package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadKeystoreByNameRequest Request Object
type DownloadKeystoreByNameRequest struct {

	// 文件名
	Name string `json:"name"`

	// 租户ID
	DomainId string `json:"domain_id"`

	// 文件ID
	Id string `json:"id"`
}

func (o DownloadKeystoreByNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadKeystoreByNameRequest struct{}"
	}

	return strings.Join([]string{"DownloadKeystoreByNameRequest", string(data)}, " ")
}
