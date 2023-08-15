package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectTemplatesRequest Request Object
type ListProjectTemplatesRequest struct {

	// cpu架构 x86|arm
	Arch *string `json:"arch,omitempty"`

	// 技术栈ID，通过技术栈管理ListStacks接口获取。
	StackId string `json:"stack_id"`
}

func (o ListProjectTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListProjectTemplatesRequest", string(data)}, " ")
}
