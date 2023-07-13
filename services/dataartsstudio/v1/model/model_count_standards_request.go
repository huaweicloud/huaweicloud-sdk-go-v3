package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountStandardsRequest Request Object
type CountStandardsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 实体id
	Id string `json:"id"`

	// 按业务类型查询
	BizType *string `json:"biz_type,omitempty"`
}

func (o CountStandardsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountStandardsRequest struct{}"
	}

	return strings.Join([]string{"CountStandardsRequest", string(data)}, " ")
}
