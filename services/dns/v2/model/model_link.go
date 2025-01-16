package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Link 指向当前资源的链接。
type Link struct {

	// 当前资源的链接。
	Self *string `json:"self,omitempty"`
}

func (o Link) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Link struct{}"
	}

	return strings.Join([]string{"Link", string(data)}, " ")
}
