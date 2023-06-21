package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPackageDetailResponse struct {

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

	// 习题库描述
	Description *string `json:"description,omitempty"`

	// 习题库里的习题数量
	ExerciseCnt    *int32 `json:"exercise_cnt,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowPackageDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPackageDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowPackageDetailResponse", string(data)}, " ")
}
