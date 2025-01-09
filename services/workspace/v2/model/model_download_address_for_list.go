package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DownloadAddressForList struct {

	// 录屏记录（文件）UUID。
	Id *string `json:"id,omitempty"`

	// 下载地址。
	DownloadUrl *string `json:"download_url,omitempty"`
}

func (o DownloadAddressForList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAddressForList struct{}"
	}

	return strings.Join([]string{"DownloadAddressForList", string(data)}, " ")
}
