package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateThumbnailsTaskResponse struct {

	// 任务ID。
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 任务状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 任务创建时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	Output *ObsObjInfo `json:"output,omitempty" xml:"output"`

	// 截图文件名称
	OutputFileName *string `json:"output_file_name,omitempty" xml:"output_file_name"`

	// 指定的截图时间点
	ThumbnailTime *string `json:"thumbnail_time,omitempty" xml:"thumbnail_time"`

	// 截图任务描述，当截图出现异常时，此字段为异常的原因
	Description    *string `json:"description,omitempty" xml:"description"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateThumbnailsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateThumbnailsTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateThumbnailsTaskResponse", string(data)}, " ")
}
