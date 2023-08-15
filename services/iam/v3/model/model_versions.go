package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Versions
type Versions struct {

	// 版本的资源链接信息。
	Values []Version `json:"values"`
}

func (o Versions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Versions struct{}"
	}

	return strings.Join([]string{"Versions", string(data)}, " ")
}
