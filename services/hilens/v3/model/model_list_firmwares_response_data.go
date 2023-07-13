package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFirmwaresResponseData 固件详情
type ListFirmwaresResponseData struct {

	// 固件id
	Id string `json:"id"`

	// 固件名称
	Name string `json:"name"`

	// 固件描述
	Description string `json:"description"`

	// 固件版本
	Version string `json:"version"`

	// 固件版本类型
	VersionType string `json:"version_type"`

	// 固件到期时间
	ExpireTime string `json:"expire_time"`

	// 固件白名单
	FirmwareWhitelist []string `json:"firmware_whitelist"`

	// 固件类型
	Type string `json:"type"`

	// 产品系列
	Series string `json:"series"`

	// 固件适用设备类型
	DeviceType string `json:"device_type"`

	// 边缘节点架构
	Arch string `json:"arch"`

	// 边缘设备操作系统名称
	OsName string `json:"os_name"`

	// 边缘节点操作系统类型
	OsType string `json:"os_type"`

	// 边缘设备操作系统版本
	OsVersion string `json:"os_version"`

	// 当前固件大小(单位:byte)
	Size int64 `json:"size"`

	// 创建时间毫秒数
	CreateTime int64 `json:"create_time"`

	// 更新时间毫秒数
	UpdateTime int64 `json:"update_time"`
}

func (o ListFirmwaresResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFirmwaresResponseData struct{}"
	}

	return strings.Join([]string{"ListFirmwaresResponseData", string(data)}, " ")
}
