package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DownloadDataTraceRequest struct {

	// Locale语言信息, zh_cn返回中文，en-us返回英文
	XLanguage string `json:"X-Language"`

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`
}

func (o DownloadDataTraceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadDataTraceRequest struct{}"
	}

	return strings.Join([]string{"DownloadDataTraceRequest", string(data)}, " ")
}
