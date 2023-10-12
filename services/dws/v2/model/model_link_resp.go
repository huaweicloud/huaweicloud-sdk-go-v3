package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LinkResp struct {

	// 关联
	Rel *string `json:"rel,omitempty"`

	// 连接
	Href *string `json:"href,omitempty"`
}

func (o LinkResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinkResp struct{}"
	}

	return strings.Join([]string{"LinkResp", string(data)}, " ")
}
