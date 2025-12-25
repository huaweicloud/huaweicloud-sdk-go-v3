package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatasetItem struct {

	// 是否触发告警
	Alert bool `json:"alert"`

	// 是否允许告警
	AllowAlert bool `json:"allow_alert"`

	// 是否允许长期存储
	AllowLts bool `json:"allow_lts"`

	// 创建时间，单位为毫秒的时间戳
	CreateTime int64 `json:"create_time"`

	// 租户ID，唯一标识
	DomainId string `json:"domain_id"`

	// 启用状态，例如 \"active\" 表示启用
	Enable string `json:"enable"`

	// 项目ID，唯一标识
	ProjectId string `json:"project_id"`

	// 是否是区域级数据
	Region bool `json:"region"`

	// 区域ID，表示当前区域
	RegionId string `json:"region_id"`

	// 操作是否成功
	Success bool `json:"success"`

	// 总记录数
	Total int32 `json:"total"`

	// 更新时间，单位为毫秒的时间戳
	UpdateTime int64 `json:"update_time"`

	// 工作空间ID，唯一标识
	WorkspaceId string `json:"workspace_id"`
}

func (o DatasetItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatasetItem struct{}"
	}

	return strings.Join([]string{"DatasetItem", string(data)}, " ")
}
