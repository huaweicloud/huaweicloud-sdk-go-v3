package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagRequestDetail struct {

	// 资源标签对列表
	Tags *[]TagObject `json:"tags,omitempty"`
}

func (o TagRequestDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagRequestDetail struct{}"
	}

	return strings.Join([]string{"TagRequestDetail", string(data)}, " ")
}
