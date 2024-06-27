package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteTagsDto struct {

	// 标签列表
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o DeleteTagsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagsDto struct{}"
	}

	return strings.Join([]string{"DeleteTagsDto", string(data)}, " ")
}
