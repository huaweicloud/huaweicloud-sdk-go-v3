package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRemoteMirrorResponse Response Object
type ShowRemoteMirrorResponse struct {

	// **参数解释：** 仓库镜像配置ID。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 仓库ID。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** 同步状态。
	UpdateStatus *string `json:"update_status,omitempty"`

	// **参数解释：** 最近修改时间。
	LastUpdateAt *string `json:"last_update_at,omitempty"`

	// **参数解释：** 镜像地址。
	Url *string `json:"url,omitempty"`

	// **参数解释：** 最近同步成功时间。
	LastSuccessfulUpdateAt *string `json:"last_successful_update_at,omitempty"`

	// **参数解释：** 同步失败次数。
	NumberOfFailures *int32 `json:"number_of_failures,omitempty"`

	// **参数解释：** 开启定时同步。
	MirroringEnabled *bool `json:"mirroring_enabled,omitempty"`

	// **参数解释：** 私有镜像。
	IsPrivate *bool `json:"is_private,omitempty"`

	// **参数解释：** 拓展点uuid。
	EndpointUuid *string `json:"endpoint_uuid,omitempty"`

	// **参数解释：** 最近失败信息。
	LastError      *string `json:"last_error,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRemoteMirrorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRemoteMirrorResponse struct{}"
	}

	return strings.Join([]string{"ShowRemoteMirrorResponse", string(data)}, " ")
}
