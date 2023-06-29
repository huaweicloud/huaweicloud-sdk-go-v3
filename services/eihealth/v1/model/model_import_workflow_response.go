package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportWorkflowResponse Response Object
type ImportWorkflowResponse struct {

	// 流程ID
	Id *string `json:"id,omitempty"`

	// 导入流程结果状态, 包括以下状态：IMPORT_SUCCESS,IMPORT_FAIL
	ImportWorkflowStatus *string `json:"import_workflow_status,omitempty"`

	// 导入应用详情
	ImportAppResults *[]ImportAppRsp `json:"import_app_results,omitempty"`
	HttpStatusCode   int             `json:"-"`
}

func (o ImportWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportWorkflowResponse struct{}"
	}

	return strings.Join([]string{"ImportWorkflowResponse", string(data)}, " ")
}
