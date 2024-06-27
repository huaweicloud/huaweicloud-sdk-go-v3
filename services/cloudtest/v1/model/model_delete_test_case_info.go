package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteTestCaseInfo struct {

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 分支uri
	VersionUri *string `json:"version_uri,omitempty"`

	// 用例id列表
	CaseIds []string `json:"case_ids"`
}

func (o DeleteTestCaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTestCaseInfo struct{}"
	}

	return strings.Join([]string{"DeleteTestCaseInfo", string(data)}, " ")
}
