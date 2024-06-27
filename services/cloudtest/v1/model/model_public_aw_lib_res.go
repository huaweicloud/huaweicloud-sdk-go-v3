package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublicAwLibRes struct {

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建时间戳
	CreateTimeStamp *int64 `json:"create_time_stamp,omitempty"`

	// 创建时间字符串
	CreateTimeString *string `json:"create_time_string,omitempty"`

	// 创建人
	CreateUser *string `json:"create_user,omitempty"`

	// 关联公共aw时的编辑链接
	DocumentLink *string `json:"document_link,omitempty"`

	// id
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 蓝区:Blue,绿区:Green,黄区:Yellow
	NetArea *[]string `json:"net_area,omitempty"`

	// obs临时url
	ObsTempUrl *string `json:"obs_temp_url,omitempty"`

	// 产品线
	ProductLine *string `json:"product_line,omitempty"`

	// 仓库分支
	RepoBranch *string `json:"repo_branch,omitempty"`

	// 存公共aw依赖jar包的目录
	RepoLibPath *string `json:"repo_lib_path,omitempty"`

	// 仓库密码
	RepoPassword *string `json:"repo_password,omitempty"`

	// 仓库秘钥
	RepoPrivateKey *string `json:"repo_private_key,omitempty"`

	// 仓库地址
	RepoUrl *string `json:"repo_url,omitempty"`

	// 仓库用户名
	RepoUsername *string `json:"repo_username,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 更新时间戳
	UpdateTimeStamp *int64 `json:"update_time_stamp,omitempty"`

	// 更新时间字符串
	UpdateTimeString *string `json:"update_time_string,omitempty"`

	// 更新人
	UpdateUser *string `json:"update_user,omitempty"`
}

func (o PublicAwLibRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicAwLibRes struct{}"
	}

	return strings.Join([]string{"PublicAwLibRes", string(data)}, " ")
}
