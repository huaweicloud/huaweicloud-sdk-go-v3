package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRequirementsOverviewRequest Request Object
type ShowRequirementsOverviewRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 版本id
	VersionId string `json:"version_id"`

	Body *QueryRequirementsOverviewInfo `json:"body,omitempty"`
}

func (o ShowRequirementsOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRequirementsOverviewRequest struct{}"
	}

	return strings.Join([]string{"ShowRequirementsOverviewRequest", string(data)}, " ")
}
