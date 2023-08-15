package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnvResponse Response Object
type CreateEnvResponse struct {

	// 对象id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEnvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvResponse struct{}"
	}

	return strings.Join([]string{"CreateEnvResponse", string(data)}, " ")
}
