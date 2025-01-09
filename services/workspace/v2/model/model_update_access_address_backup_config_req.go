package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessAddressBackupConfigReq 修改云办公服务接入地址备份配置请求体。
type UpdateAccessAddressBackupConfigReq struct {

	// 接入配置列表信息。
	AccessConfig *[]AccessConfigInfo `json:"access_config,omitempty"`
}

func (o UpdateAccessAddressBackupConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessAddressBackupConfigReq struct{}"
	}

	return strings.Join([]string{"UpdateAccessAddressBackupConfigReq", string(data)}, " ")
}
