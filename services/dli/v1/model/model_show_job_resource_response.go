package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobResourceResponse Response Object
type ShowJobResourceResponse struct {

	// 资源包上传的unix时间。是单位为“毫秒”的时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新已上传资源包的unix时间。是单位为“毫秒”的时间戳。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 资源类型。
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源名。
	ResourceName *string `json:"resource_name,omitempty"`

	// \"UPLOADING\"表示正在上传 \"READY\"表示资源包已上传 \"FAILED\"表示资源包上传失败
	Status *string `json:"status,omitempty"`

	// 资源包在队列中的名字。
	UnderlyingName *string `json:"underlying_name,omitempty"`

	// 资源包拥有者
	Owner          *string `json:"owner,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobResourceResponse struct{}"
	}

	return strings.Join([]string{"ShowJobResourceResponse", string(data)}, " ")
}
