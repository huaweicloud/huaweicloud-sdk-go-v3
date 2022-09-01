package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 模块
type WorkitemsModule struct {

	// 模块id
	Id *string `json:"id,omitempty" xml:"id"`

	// 模块
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o WorkitemsModule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkitemsModule struct{}"
	}

	return strings.Join([]string{"WorkitemsModule", string(data)}, " ")
}
