package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateConfigMapResponse struct {
	Configmap *ConfigMap `json:"configmap,omitempty"`

	// 工作空间ID，默认为注册账号/子账号的default工作空间，可通过专业版HiLens控制台展开工作空间列表获取到工作空间ID
	WorkspaceId    *string `json:"workspace_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateConfigMapResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigMapResponse struct{}"
	}

	return strings.Join([]string{"UpdateConfigMapResponse", string(data)}, " ")
}
