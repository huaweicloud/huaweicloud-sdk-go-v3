package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkTableIssuseListResponseBodyProject 项目信息
type WorkTableIssuseListResponseBodyProject struct {

	// 项目id
	Id *int32 `json:"id,omitempty"`

	// 项目uuid
	Identifier *string `json:"identifier,omitempty"`

	// 项目名称
	Name *string `json:"name,omitempty"`

	// 项目类型
	Type *string `json:"type,omitempty"`
}

func (o WorkTableIssuseListResponseBodyProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkTableIssuseListResponseBodyProject struct{}"
	}

	return strings.Join([]string{"WorkTableIssuseListResponseBodyProject", string(data)}, " ")
}
