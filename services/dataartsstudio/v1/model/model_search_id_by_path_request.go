package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchIdByPathRequest Request Object
type SearchIdByPathRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 路径
	Path string `json:"path"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`
}

func (o SearchIdByPathRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchIdByPathRequest struct{}"
	}

	return strings.Join([]string{"SearchIdByPathRequest", string(data)}, " ")
}
