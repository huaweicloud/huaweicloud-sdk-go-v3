package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStylesRequest Request Object
type ListStylesRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 开发者应用作为资产权属的可选字段。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 任务状态，默认所有状态。  可多个状态查询，使用英文逗号分隔。  如state=CREATING,PUBLISHED
	State *string `json:"state,omitempty"`

	// 排序字段，目前只支持create_time。
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方式。 * asc：升序 * desc：降序  默认asc升序。
	SortDir *string `json:"sort_dir,omitempty"`

	// 过滤创建时间<=输入时间的记录。
	CreateUntil *string `json:"create_until,omitempty"`

	// 过滤创建时间>=输入时间的记录。
	CreateSince *string `json:"create_since,omitempty"`
}

func (o ListStylesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStylesRequest struct{}"
	}

	return strings.Join([]string{"ListStylesRequest", string(data)}, " ")
}
