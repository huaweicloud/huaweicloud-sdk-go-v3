package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportImageResponse Response Object
type ImportImageResponse struct {

	// 镜像id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportImageResponse struct{}"
	}

	return strings.Join([]string{"ImportImageResponse", string(data)}, " ")
}
