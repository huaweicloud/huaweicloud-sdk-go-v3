package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogDownloadRequest Request Object
type ShowSlowLogDownloadRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`
}

func (o ShowSlowLogDownloadRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogDownloadRequest struct{}"
	}

	return strings.Join([]string{"ShowSlowLogDownloadRequest", string(data)}, " ")
}
