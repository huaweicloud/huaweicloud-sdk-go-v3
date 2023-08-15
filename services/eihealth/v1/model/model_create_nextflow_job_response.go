package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNextflowJobResponse Response Object
type CreateNextflowJobResponse struct {

	// 作业id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateNextflowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNextflowJobResponse struct{}"
	}

	return strings.Join([]string{"CreateNextflowJobResponse", string(data)}, " ")
}
