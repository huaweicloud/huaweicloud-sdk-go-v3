package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowProjectTraceDataResponse struct {

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
	DownloadUrl    *string `json:"download_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowProjectTraceDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectTraceDataResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectTraceDataResponse", string(data)}, " ")
}
