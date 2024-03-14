package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuerySupportNodeTypeBean 规格信息
type QuerySupportNodeTypeBean struct {

	// 规格类型
	NodeType *string `json:"node_type,omitempty"`

	// 是否售罄
	IsSellout *bool `json:"is_sellout,omitempty"`
}

func (o QuerySupportNodeTypeBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySupportNodeTypeBean struct{}"
	}

	return strings.Join([]string{"QuerySupportNodeTypeBean", string(data)}, " ")
}
