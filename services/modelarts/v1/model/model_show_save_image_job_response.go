package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSaveImageJobResponse Response Object
type ShowSaveImageJobResponse struct {

	// 镜像名称，长度限制512个字符，支持小写字母、数字、中划线、下划线和点。
	Name *string `json:"name,omitempty"`

	// 镜像所属组织，可以在SWR控制台“组织管理”创建和查看。
	Namespace *string `json:"namespace,omitempty"`

	// 镜像tag，长度限制64个字符， 支持大小写字母、数字、中划线、下划线和点。
	Tag *string `json:"tag,omitempty"`

	// 该镜像所对应的描述信息，长度限制512个字符。
	Description *string `json:"description,omitempty"`

	// 镜像状态。枚举值如下： - INIT：初始化。 - CREATING：镜像保存中，此时训练作业不可用。 - CREATE_FAILED：镜像保存失败。 - ACTIVE：镜像保存成功，保存的镜像可以在SWR控制台查看，同时可以基于保存的镜像创建训练作业。
	Status *string `json:"status,omitempty"`

	// 镜像创建的时间，UTC毫秒。
	Message *string `json:"message,omitempty"`

	// 镜像保存操作过程中，展示构建信息。
	CreateTime     *int64 `json:"create_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSaveImageJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSaveImageJobResponse struct{}"
	}

	return strings.Join([]string{"ShowSaveImageJobResponse", string(data)}, " ")
}
