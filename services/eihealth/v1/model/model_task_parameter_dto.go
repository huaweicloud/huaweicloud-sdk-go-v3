package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskParameterDto struct {

	// 子任务的参数名称，长度为[1,32]，以小写字母开头，允许出现中划线(-)、小写字母和数字，且必须以小写字母或数字结尾。需要和已有应用的参数名称一致。
	Name string `json:"name"`

	// 子任务的参数类型，不填默认为MANUAL
	Source *string `json:"source,omitempty"`

	// 子任务的参数数值，根据参数类型进行合法性校验
	Values *[]string `json:"values,omitempty"`
}

func (o TaskParameterDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskParameterDto struct{}"
	}

	return strings.Join([]string{"TaskParameterDto", string(data)}, " ")
}
