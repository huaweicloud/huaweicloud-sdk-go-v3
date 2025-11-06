package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TreeObjectDto struct {

	// 文件唯一标识
	Id *string `json:"id,omitempty"`

	// 文件名称
	Name *string `json:"name,omitempty"`

	// 对象类型
	Type *string `json:"type,omitempty"`

	// 文件路径
	Path *string `json:"path,omitempty"`

	// 模式结构
	Mode *string `json:"mode,omitempty"`

	// 子模块链接
	SubmoduleLink *string `json:"submodule_link,omitempty"`

	// 子模块分支
	SubmoduleBranch *string `json:"submodule_branch,omitempty"`

	// md5值
	Md5 *string `json:"md5,omitempty"`
}

func (o TreeObjectDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TreeObjectDto struct{}"
	}

	return strings.Join([]string{"TreeObjectDto", string(data)}, " ")
}
