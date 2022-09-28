package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据对象
type TraceDataRsp struct {

	// 对象全路径（项目名称:/路径）
	Path *string `json:"path,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	Type *PathType `json:"type,omitempty"`

	// 大小
	Size *int64 `json:"size,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 下载链接
	DownloadUrl *string `json:"download_url,omitempty"`
}

func (o TraceDataRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TraceDataRsp struct{}"
	}

	return strings.Join([]string{"TraceDataRsp", string(data)}, " ")
}
