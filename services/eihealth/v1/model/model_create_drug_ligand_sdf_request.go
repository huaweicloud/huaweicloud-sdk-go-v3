package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrugLigandSdfRequest Request Object
type CreateDrugLigandSdfRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreateLigandSdfReq `json:"body,omitempty"`
}

func (o CreateDrugLigandSdfRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugLigandSdfRequest struct{}"
	}

	return strings.Join([]string{"CreateDrugLigandSdfRequest", string(data)}, " ")
}
