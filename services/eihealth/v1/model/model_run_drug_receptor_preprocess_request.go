package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunDrugReceptorPreprocessRequest Request Object
type RunDrugReceptorPreprocessRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
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
