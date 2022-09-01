package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAnimatedGraphicsTaskReq struct {
	Input *ObsObjInfo `json:"input,omitempty" xml:"input"`

	Output *ObsObjInfo `json:"output,omitempty" xml:"output"`

	// 用户自定义数据。
	UserData *string `json:"user_data,omitempty" xml:"user_data"`

	OutputParam *AnimatedGraphicsOutputParam `json:"output_param,omitempty" xml:"output_param"`
}

func (o CreateAnimatedGraphicsTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnimatedGraphicsTaskReq struct{}"
	}

	return strings.Join([]string{"CreateAnimatedGraphicsTaskReq", string(data)}, " ")
}
