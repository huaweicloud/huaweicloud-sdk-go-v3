package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageResponse Response Object
type CreateImageResponse struct {

	// 镜像id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageResponse struct{}"
	}

	return strings.Join([]string{"CreateImageResponse", string(data)}, " ")
}
