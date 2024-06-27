package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddTestCaseCommentResponse Response Object
type AddTestCaseCommentResponse struct {
	Uri *string `json:"uri,omitempty"`

	Creator *string `json:"creator,omitempty"`

	Comment *string `json:"comment,omitempty"`

	Notifier *[]string `json:"notifier,omitempty"`

	TestCaseUri *string `json:"test_case_uri,omitempty"`

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 创建时间时间戳
	CreateTimeTimestamp *int64 `json:"create_time_timestamp,omitempty"`

	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 更新时间时间戳
	UpdateTimeTimestamp *int64 `json:"update_time_timestamp,omitempty"`

	ProjectUuid *string `json:"project_uuid,omitempty"`

	VersionUri *string `json:"version_uri,omitempty"`

	DisplayName    *string `json:"display_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddTestCaseCommentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTestCaseCommentResponse struct{}"
	}

	return strings.Join([]string{"AddTestCaseCommentResponse", string(data)}, " ")
}
