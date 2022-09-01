package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateExtractTaskReq struct {
	Input *ObsObjInfo `json:"input,omitempty" xml:"input"`

	Output *ObsObjInfo `json:"output,omitempty" xml:"output"`

	// 用户自定义数据。
	UserData *string `json:"user_data,omitempty" xml:"user_data"`

	// 是否同步处理, - 0：排队处理 - 1：同步处理  默认值：0
	Sync *int32 `json:"sync,omitempty" xml:"sync"`
}

func (o CreateExtractTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExtractTaskReq struct{}"
	}

	return strings.Join([]string{"CreateExtractTaskReq", string(data)}, " ")
}
