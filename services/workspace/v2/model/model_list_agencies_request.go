package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgenciesRequest Request Object
type ListAgenciesRequest struct {

	// 委托场景，多个用英文逗号分隔。 - WORKSPACE：云桌面。 - CLOUD_GAME：云游戏。 - CLOUD_STORAGE 云存储。 - SCREEN_RECORD：录屏审计。
	Scene *string `json:"scene,omitempty"`

	// 操作类型。 - CREATE 创建 - FIX 修复
	Action *string `json:"action,omitempty"`
}

func (o ListAgenciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgenciesRequest struct{}"
	}

	return strings.Join([]string{"ListAgenciesRequest", string(data)}, " ")
}
