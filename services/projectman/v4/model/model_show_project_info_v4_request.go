package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowProjectInfoV4Request struct {
	// devcloud的项目ID

	ProjectId string `json:"project_id"`
}

func (o ShowProjectInfoV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectInfoV4Request struct{}"
	}

	return strings.Join([]string{"ShowProjectInfoV4Request", string(data)}, " ")
}
