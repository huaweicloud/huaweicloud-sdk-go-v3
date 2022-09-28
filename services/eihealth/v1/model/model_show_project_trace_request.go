package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowProjectTraceRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 指定文件夹（路径）
	Path string `json:"path"`

	// 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowProjectTraceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectTraceRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectTraceRequest", string(data)}, " ")
}
