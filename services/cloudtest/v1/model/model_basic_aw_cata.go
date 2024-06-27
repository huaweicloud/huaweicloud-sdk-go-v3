package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasicAwCata struct {

	// aw归属目录信息
	AwDir *string `json:"aw_dir,omitempty"`

	// 目录层级
	CataType *int32 `json:"cata_type,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建人
	CreateUser *string `json:"create_user,omitempty"`

	// 创建人id
	CreateUserId *string `json:"create_user_id,omitempty"`

	// 目录描述
	Desc *string `json:"desc,omitempty"`

	// 引用次数
	ExtraInfo *string `json:"extra_info,omitempty"`

	// id
	Id *string `json:"id,omitempty"`

	// 判断是否为文件夹的标识
	IsFolder *string `json:"is_folder,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// aw在页面上显示的名字
	NameView *string `json:"nameView,omitempty"`

	// aw目录父编号
	ParentId *string `json:"parent_id,omitempty"`

	// 工程ID
	ProjectId *string `json:"project_id,omitempty"`

	// 引用次数
	RefCnt *int32 `json:"refCnt,omitempty"`

	// 区域名称
	Region *string `json:"region,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 更新人
	UpdateUser *string `json:"update_user,omitempty"`
}

func (o BasicAwCata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicAwCata struct{}"
	}

	return strings.Join([]string{"BasicAwCata", string(data)}, " ")
}
