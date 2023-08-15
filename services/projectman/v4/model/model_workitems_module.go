package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkitemsModule 模块
type WorkitemsModule struct {

	// 模块id
	Id *string `json:"id,omitempty"`

	// 模块
	Name *string `json:"name,omitempty"`
}

func (o WorkitemsModule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkitemsModule struct{}"
	}

	return strings.Join([]string{"WorkitemsModule", string(data)}, " ")
}
