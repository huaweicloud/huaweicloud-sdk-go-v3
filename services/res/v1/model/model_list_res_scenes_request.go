package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResScenesRequest Request Object
type ListResScenesRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	// 场景类型： - customize，自定义推荐 - intelligent，智能场景
	Category string `json:"category"`
}

func (o ListResScenesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResScenesRequest struct{}"
	}

	return strings.Join([]string{"ListResScenesRequest", string(data)}, " ")
}
