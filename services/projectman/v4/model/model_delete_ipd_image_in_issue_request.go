package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIpdImageInIssueRequest Request Object
type DeleteIpdImageInIssueRequest struct {

	// 项目Id
	ProjectId string `json:"project_id"`

	// 图片绑定的工作项id
	IssueId string `json:"issue_id"`

	// 工作项描述中的图片名称,如f401ab8826b1426285910e6c5da80d07.png
	FileName string `json:"file_name"`
}

func (o DeleteIpdImageInIssueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpdImageInIssueRequest struct{}"
	}

	return strings.Join([]string{"DeleteIpdImageInIssueRequest", string(data)}, " ")
}
