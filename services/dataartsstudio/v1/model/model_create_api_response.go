package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApiResponse Response Object
type CreateApiResponse struct {

	// 创建成功的API ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApiResponse struct{}"
	}

	return strings.Join([]string{"CreateApiResponse", string(data)}, " ")
}
