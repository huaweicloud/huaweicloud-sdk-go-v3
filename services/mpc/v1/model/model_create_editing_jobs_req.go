package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建剪辑任务
type CreateEditingJobsReq struct {

	// 指定片源多个剪切时间段
	Inputs *[]EditingInput `json:"inputs,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	EditingSettings *EditingSetting `json:"editing_settings,omitempty"`

	// 用户自定义数据
	UserData *string `json:"user_data,omitempty"`
}

func (o CreateEditingJobsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEditingJobsReq struct{}"
	}

	return strings.Join([]string{"CreateEditingJobsReq", string(data)}, " ")
}
