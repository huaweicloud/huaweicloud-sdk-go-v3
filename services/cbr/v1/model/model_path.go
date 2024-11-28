package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Path
type Path struct {

	// 路径ID
	Id *string `json:"id,omitempty"`

	// 路径状态，有available和remove两种状态
	Status *string `json:"status,omitempty"`

	// 该路径所属于的客户端ID
	AgentId *string `json:"agent_id,omitempty"`

	// 路径详情
	DirPath *string `json:"dir_path,omitempty"`

	// 排除目录列表，多个路径之间以英文逗号分隔 > 该特性目前处于公测阶段，部分region可能无法使用。
	ExcludePaths *string `json:"exclude_paths,omitempty"`
}

func (o Path) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Path struct{}"
	}

	return strings.Join([]string{"Path", string(data)}, " ")
}
