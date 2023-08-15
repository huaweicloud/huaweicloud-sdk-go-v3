package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoJobRequest Request Object
type ListAutoJobRequest struct {

	// Locale语言类型，zh_cn返回中文，en_us返回英文
	XLanguage *string `json:"X-Language,omitempty"`

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]
	Offset *int32 `json:"offset,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方向，升序或降序，即ASC 和DESC
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListAutoJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoJobRequest struct{}"
	}

	return strings.Join([]string{"ListAutoJobRequest", string(data)}, " ")
}
