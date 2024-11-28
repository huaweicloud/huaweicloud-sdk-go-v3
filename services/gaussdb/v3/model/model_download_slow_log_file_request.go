package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadSlowLogFileRequest Request Object
type DownloadSlowLogFileRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 节点ID。
	NodeId string `json:"node_id"`
}

func (o DownloadSlowLogFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadSlowLogFileRequest struct{}"
	}

	return strings.Join([]string{"DownloadSlowLogFileRequest", string(data)}, " ")
}
