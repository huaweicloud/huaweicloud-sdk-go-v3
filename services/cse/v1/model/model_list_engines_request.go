package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEnginesRequest struct {

	// 如果不带则默认企业项目为\"default\"，ID为\"0\"。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *string `json:"limit,omitempty"`
}

func (o ListEnginesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnginesRequest struct{}"
	}

	return strings.Join([]string{"ListEnginesRequest", string(data)}, " ")
}
