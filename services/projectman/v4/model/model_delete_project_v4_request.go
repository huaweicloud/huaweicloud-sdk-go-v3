package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteProjectV4Request struct {
	// devcloud的项目id

	ProjectId string `json:"project_id"`
}

func (o DeleteProjectV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProjectV4Request struct{}"
	}

	return strings.Join([]string{"DeleteProjectV4Request", string(data)}, " ")
}
