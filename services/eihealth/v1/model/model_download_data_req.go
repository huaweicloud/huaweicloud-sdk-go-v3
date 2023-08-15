package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadDataReq 下载文件请求体
type DownloadDataReq struct {
	Type *DownloadDataTypeEnum `json:"type,omitempty"`

	// 资源地址(支持https、obs地址)
	Url *string `json:"url,omitempty"`
}

func (o DownloadDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadDataReq struct{}"
	}

	return strings.Join([]string{"DownloadDataReq", string(data)}, " ")
}
