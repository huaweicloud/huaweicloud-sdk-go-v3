package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelGroupResourceItemResp 模型分组关联的应用对象详情。
type ModelGroupResourceItemResp struct {

	// 关联记录id。
	Id *int64 `json:"id,omitempty"`

	// 资源类型（DESKTOP-桌面实例，DESKTOP_TAG-桌面标签）。
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源id（Agent实例id或桌面标签id）。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称（桌面实例名称或桌面标签key:value格式）。
	ResourceName *string `json:"resource_name,omitempty"`

	// 关联创建时间（ISO8601格式，UTC时区）。
	CreatedTime *string `json:"created_time,omitempty"`
}

func (o ModelGroupResourceItemResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelGroupResourceItemResp struct{}"
	}

	return strings.Join([]string{"ModelGroupResourceItemResp", string(data)}, " ")
}
