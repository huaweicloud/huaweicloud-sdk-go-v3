package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppResponse Response Object
type ListAppResponse struct {

	// 是否还有更多满足条件的App。  - true：是。 - false：否。
	HasMoreApp *bool `json:"has_more_app,omitempty"`

	// AppEntry list that meets the current request.
	Apps *[]DescribeAppResult `json:"apps,omitempty"`

	// 满足条件的App总数。
	TotalNumber    *int32 `json:"total_number,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppResponse struct{}"
	}

	return strings.Join([]string{"ListAppResponse", string(data)}, " ")
}
