package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestCaseReviewsRequest Request Object
type ShowTestCaseReviewsRequest struct {

	// 项目id
	ProjectUuid string `json:"project_uuid"`

	// 版本URI
	VersionUri string `json:"version_uri"`

	// 当前页数
	PageNo int32 `json:"page_no"`

	// 每页条数
	PageSize int32 `json:"page_size"`

	// 分支用例URI
	TestcaseUri string `json:"testcase_uri"`
}

func (o ShowTestCaseReviewsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestCaseReviewsRequest struct{}"
	}

	return strings.Join([]string{"ShowTestCaseReviewsRequest", string(data)}, " ")
}
