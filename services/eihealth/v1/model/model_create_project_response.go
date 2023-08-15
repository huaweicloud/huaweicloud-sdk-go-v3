package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectResponse Response Object
type CreateProjectResponse struct {

	// 项目id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectResponse struct{}"
	}

	return strings.Join([]string{"CreateProjectResponse", string(data)}, " ")
}
