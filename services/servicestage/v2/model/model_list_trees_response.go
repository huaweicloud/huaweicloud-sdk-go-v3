package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTreesResponse struct {
	// 仓库文件列表。

	Paths          *[]string `json:"paths,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTreesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTreesResponse struct{}"
	}

	return strings.Join([]string{"ListTreesResponse", string(data)}, " ")
}
