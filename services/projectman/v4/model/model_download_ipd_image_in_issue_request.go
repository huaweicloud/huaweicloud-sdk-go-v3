package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadIpdImageInIssueRequest Request Object
type DownloadIpdImageInIssueRequest struct {

	// 项目Id
	ProjectId string `json:"project_id"`

	// 图片绑定的工作项id
	IssueId string `json:"issue_id"`

	// 工作项描述中的图片名称,如f401ab8826b1426285910e6c5da80d07.png
	FileName string `json:"file_name"`
}

func (o DownloadIpdImageInIssueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadIpdImageInIssueRequest struct{}"
	}

	return strings.Join([]string{"DownloadIpdImageInIssueRequest", string(data)}, " ")
}
