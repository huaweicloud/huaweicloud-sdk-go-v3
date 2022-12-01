package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateWorkSpaceResponse struct {

	// 创建的工作空间名
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWorkSpaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkSpaceResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkSpaceResponse", string(data)}, " ")
}
