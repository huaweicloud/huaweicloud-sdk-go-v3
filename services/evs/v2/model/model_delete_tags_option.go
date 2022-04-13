package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteTagsOption struct {
	// 标签键。

	Key string `json:"key"`
}

func (o DeleteTagsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagsOption struct{}"
	}

	return strings.Join([]string{"DeleteTagsOption", string(data)}, " ")
}
