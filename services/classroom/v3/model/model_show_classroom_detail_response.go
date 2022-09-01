package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowClassroomDetailResponse struct {

	// 课堂名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 课堂描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 课堂公告
	Announcement *string `json:"announcement,omitempty" xml:"announcement"`

	// 课堂公告创建时间，日期格式：yyyy-MM-dd
	AnnouncementTime *string `json:"announcement_time,omitempty" xml:"announcement_time"`

	// 课堂创建时间，日期格式：yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 课堂最新更新时间，日期格式：yyyy-MM-dd HH:mm:ss
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`

	// 当前课堂的授课人
	Teacher *string `json:"teacher,omitempty" xml:"teacher"`

	// 课堂学分
	Credit float32 `json:"credit,omitempty" xml:"credit"`

	// 课堂开始时间，日期格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 课堂结束时间，日期格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 当前用户在课堂下角色，取值范围：teacher：老师，student：学生
	Role *string `json:"role,omitempty" xml:"role"`

	// 授课学校
	School *string `json:"school,omitempty" xml:"school"`

	// 课堂下目录数量
	ContentCount *int32 `json:"content_count,omitempty" xml:"content_count"`

	// 课堂下课件数量
	CoursewareCount *int32 `json:"courseware_count,omitempty" xml:"courseware_count"`

	// 课堂下作业数量
	JobCount *int32 `json:"job_count,omitempty" xml:"job_count"`

	// 课堂下成员数量
	MemberCount *int32 `json:"member_count,omitempty" xml:"member_count"`

	// 课堂当前的状态，normal：课堂处于正常状态，archive：课堂已归档
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowClassroomDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClassroomDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowClassroomDetailResponse", string(data)}, " ")
}
