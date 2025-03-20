package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagResourceV5Response Response Object
type TagResourceV5Response struct {

	// 自定义标签列表。
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o TagResourceV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResourceV5Response struct{}"
	}

	return strings.Join([]string{"TagResourceV5Response", string(data)}, " ")
}
