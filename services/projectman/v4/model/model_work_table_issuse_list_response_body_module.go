package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkTableIssuseListResponseBodyModule 模块信息
type WorkTableIssuseListResponseBodyModule struct {

	// 模块id
	Id *int32 `json:"id,omitempty"`

	// 模块名称
	Name *string `json:"name,omitempty"`

	// 模块路径名称
	PathName *string `json:"path_name,omitempty"`
}

func (o WorkTableIssuseListResponseBodyModule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkTableIssuseListResponseBodyModule struct{}"
	}

	return strings.Join([]string{"WorkTableIssuseListResponseBodyModule", string(data)}, " ")
}
