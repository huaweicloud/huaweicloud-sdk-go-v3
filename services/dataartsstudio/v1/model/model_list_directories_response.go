package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDirectoriesResponse Response Object
type ListDirectoriesResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListDirectoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDirectoriesResponse struct{}"
	}

	return strings.Join([]string{"ListDirectoriesResponse", string(data)}, " ")
}
