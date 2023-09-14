package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditLogDownloadLinkRequest Request Object
type ListAuditLogDownloadLinkRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 节点ID。 - 若输入，则只获取该节点的全量SQL下载链接。 - 若不输入，则获取该实例所有节点的全量SQL下载链接。
	NodeId *string `json:"node_id,omitempty"`

	// 上次查询的最后一个文件的文件名。 - 若输入，则从该文件名以后按字典顺序开始查询。 - 若不输入，则从第一个文件开始查询。
	LastFileName *string `json:"last_file_name,omitempty"`

	// 一次查询返回的文件数量。  默认值为10，取值范围：1~50之间的整数值。
	Limit *int32 `json:"limit,omitempty"`

	// 开始时间，不得早于实例创建时间。格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	StartTime string `json:"start_time"`

	// 结束时间，不得晚于当前时间。格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	EndTime string `json:"end_time"`
}

func (o ListAuditLogDownloadLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditLogDownloadLinkRequest struct{}"
	}

	return strings.Join([]string{"ListAuditLogDownloadLinkRequest", string(data)}, " ")
}
