package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestCasesChangeStatisticsRequest Request Object
type ShowTestCasesChangeStatisticsRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 版本uri
	VersionId string `json:"version_id"`
}

func (o ShowTestCasesChangeStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestCasesChangeStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowTestCasesChangeStatisticsRequest", string(data)}, " ")
}
