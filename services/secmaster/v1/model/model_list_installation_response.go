package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstallationResponse Response Object
type ListInstallationResponse struct {

	// 计数
	Count *int64 `json:"count,omitempty"`

	// 安装脚本
	Records        *[]InstallationScript `json:"records,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListInstallationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstallationResponse struct{}"
	}

	return strings.Join([]string{"ListInstallationResponse", string(data)}, " ")
}
