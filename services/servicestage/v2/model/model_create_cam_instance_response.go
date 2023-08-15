package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCamInstanceResponse Response Object
type CreateCamInstanceResponse struct {

	// 实例ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCamInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCamInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateCamInstanceResponse", string(data)}, " ")
}
