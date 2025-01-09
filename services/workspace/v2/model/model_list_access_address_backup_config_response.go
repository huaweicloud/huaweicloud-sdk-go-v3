package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessAddressBackupConfigResponse Response Object
type ListAccessAddressBackupConfigResponse struct {

	// 接入配置列表信息。
	AccessConfig   *[]AccessConfigInfo `json:"access_config,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListAccessAddressBackupConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessAddressBackupConfigResponse struct{}"
	}

	return strings.Join([]string{"ListAccessAddressBackupConfigResponse", string(data)}, " ")
}
