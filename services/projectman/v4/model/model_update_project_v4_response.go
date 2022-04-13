package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateProjectV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateProjectV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectV4Response struct{}"
	}

	return strings.Join([]string{"UpdateProjectV4Response", string(data)}, " ")
}
