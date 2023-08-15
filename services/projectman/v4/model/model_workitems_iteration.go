package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkitemsIteration 迭代
type WorkitemsIteration struct {

	// 迭代id
	Id *string `json:"id,omitempty"`

	// 迭代名
	Name *string `json:"name,omitempty"`
}

func (o WorkitemsIteration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkitemsIteration struct{}"
	}

	return strings.Join([]string{"WorkitemsIteration", string(data)}, " ")
}
