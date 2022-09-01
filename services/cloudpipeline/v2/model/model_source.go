package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 源码仓库参数
type Source struct {

	// 源码仓名字
	CodehubName string `json:"codehub_name" xml:"codehub_name"`

	// 触发分支
	Branches []string `json:"branches" xml:"branches"`

	// 源码仓来源
	ScmType string `json:"scm_type" xml:"scm_type"`

	// 是否开启触发执行流水线功能
	HookFlag bool `json:"hook_flag" xml:"hook_flag"`

	// 触发分支
	Branch string `json:"branch" xml:"branch"`

	// 源码仓ssh地址
	GitUrl string `json:"git_url" xml:"git_url"`

	// 源码仓ID
	CodehubId string `json:"codehub_id" xml:"codehub_id"`

	// 源码仓首页url
	WebUrl string `json:"web_url" xml:"web_url"`

	// 分支列表
	BranchList []string `json:"branch_list" xml:"branch_list"`

	// 初始化ID
	InitId string `json:"init_id" xml:"init_id"`

	// 是否废弃
	Disable bool `json:"disable" xml:"disable"`
}

func (o Source) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Source struct{}"
	}

	return strings.Join([]string{"Source", string(data)}, " ")
}
