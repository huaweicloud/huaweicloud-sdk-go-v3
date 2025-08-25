package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessKeyInfo struct {

	// 访问密钥ID
	AccessKeyId string `json:"access_key_id"`

	// 访问密钥名称
	KeyName string `json:"key_name"`

	// 访问密钥AK
	AccessKey string `json:"access_key"`

	// 访问密钥所属的应用名称
	AppName string `json:"app_name"`

	// 访问密钥状态
	Status string `json:"status"`

	// 应用的创建时间，UNIX时间戳，单位毫秒
	CreateTime int64 `json:"create_time"`

	// 下载时间
	DownloadTime *int64 `json:"download_time,omitempty"`

	// 是否下载
	IsDownloaded bool `json:"is_downloaded"`

	// 是否导入
	IsImported bool `json:"is_imported"`
}

func (o AccessKeyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessKeyInfo struct{}"
	}

	return strings.Join([]string{"AccessKeyInfo", string(data)}, " ")
}
