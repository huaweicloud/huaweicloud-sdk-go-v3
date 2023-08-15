package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunDrugReceptorPreprocessRequest Request Object
type RunDrugReceptorPreprocessRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *RunReceptorPreprocessReq `json:"body,omitempty"`
}

func (o RunDrugReceptorPreprocessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDrugReceptorPreprocessRequest struct{}"
	}

	return strings.Join([]string{"RunDrugReceptorPreprocessRequest", string(data)}, " ")
}
