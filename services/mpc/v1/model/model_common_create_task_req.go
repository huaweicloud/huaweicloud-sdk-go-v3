package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommonCreateTaskReq struct {
	Input *ObsObjInfo `json:"input,omitempty" xml:"input"`

	Output *ObsObjInfo `json:"output,omitempty" xml:"output"`

	// 用户自定义数据。
	UserData *string `json:"user_data,omitempty" xml:"user_data"`
}

func (o CommonCreateTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonCreateTaskReq struct{}"
	}

	return strings.Join([]string{"CommonCreateTaskReq", string(data)}, " ")
}
