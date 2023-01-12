package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDataPathPolicyRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 对象全路径（项目名称:|路径）
	Path string `json:"path"`

	Body *DataPolicyReq `json:"body,omitempty"`
}

func (o UpdateDataPathPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataPathPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateDataPathPolicyRequest", string(data)}, " ")
}
