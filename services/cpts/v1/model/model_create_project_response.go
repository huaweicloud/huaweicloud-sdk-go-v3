package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateProjectResponse struct {

	// code
	Code *string `json:"code,omitempty" xml:"code"`

	// message
	Message *string `json:"message,omitempty" xml:"message"`

	// project_id
	ProjectId      *int32 `json:"project_id,omitempty" xml:"project_id"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectResponse struct{}"
	}

	return strings.Join([]string{"CreateProjectResponse", string(data)}, " ")
}
