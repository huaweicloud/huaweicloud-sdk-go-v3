package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataResponse Response Object
type ShowDataResponse struct {

	// 对象全路径（项目名称:/路径）
	Path *string `json:"path,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	Type *PathType `json:"type,omitempty"`

	// 大小
	Size *int64 `json:"size,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 文件内容
	Content *string `json:"content,omitempty"`

	// 下载链接
	DownloadUrl *string `json:"download_url,omitempty"`

	// 可操作标记
	AllowedOperate *bool `json:"allowed_operate,omitempty"`

	// 可删除标记
	Deletable      *bool `json:"deletable,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataResponse struct{}"
	}

	return strings.Join([]string{"ShowDataResponse", string(data)}, " ")
}
