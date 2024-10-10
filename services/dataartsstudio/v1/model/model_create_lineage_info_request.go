package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLineageInfoRequest Request Object
type CreateLineageInfoRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)
	Workspace string `json:"workspace"`

	Body *LineageInfoRequest `json:"body,omitempty"`
}

func (o CreateLineageInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLineageInfoRequest struct{}"
	}

	return strings.Join([]string{"CreateLineageInfoRequest", string(data)}, " ")
}
