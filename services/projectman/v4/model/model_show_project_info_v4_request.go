package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowProjectInfoV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`
}

func (o ShowProjectInfoV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectInfoV4Request struct{}"
	}

	return strings.Join([]string{"ShowProjectInfoV4Request", string(data)}, " ")
}
