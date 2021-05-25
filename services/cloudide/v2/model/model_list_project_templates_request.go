package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProjectTemplatesRequest struct {
	// cpu架构 x86|arm

	Arch *string `json:"arch,omitempty"`
	// 技术栈ID，通过技术栈管理ListStacksByTag接口获取。

	StackId string `json:"stack_id"`
}

func (o ListProjectTemplatesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProjectTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListProjectTemplatesRequest", string(data)}, " ")
}
