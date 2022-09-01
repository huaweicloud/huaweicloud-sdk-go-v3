package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 迭代
type WorkitemsIteration struct {

	// 迭代id
	Id *string `json:"id,omitempty" xml:"id"`

	// 迭代名
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o WorkitemsIteration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkitemsIteration struct{}"
	}

	return strings.Join([]string{"WorkitemsIteration", string(data)}, " ")
}
