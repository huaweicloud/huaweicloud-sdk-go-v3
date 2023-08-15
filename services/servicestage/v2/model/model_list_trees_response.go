package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTreesResponse Response Object
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
