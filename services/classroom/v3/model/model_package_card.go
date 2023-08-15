package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PackageCard 习题库简单信息返回体，用于列表返回
type PackageCard struct {

	// 习题库id
	Id *string `json:"id,omitempty"`

	// 习题库名称
	Name *string `json:"name,omitempty"`

	// 标签名称
	TagName *string `json:"tag_name,omitempty"`

	// 学习名称
	School *string `json:"school,omitempty"`

	// 教师名称
	TeacherName *string `json:"teacher_name,omitempty"`

	// 租户习题库编号
	OrderCount *int32 `json:"order_count,omitempty"`

	// 背景图url
	ImageUrl *string `json:"image_url,omitempty"`
}

func (o PackageCard) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackageCard struct{}"
	}

	return strings.Join([]string{"PackageCard", string(data)}, " ")
}
