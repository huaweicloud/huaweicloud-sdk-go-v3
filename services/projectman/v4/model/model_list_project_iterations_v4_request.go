package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListProjectIterationsV4Request struct {
	// devcloud的项目id

	ProjectId string `json:"project_id"`
}

func (o ListProjectIterationsV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectIterationsV4Request struct{}"
	}

	return strings.Join([]string{"ListProjectIterationsV4Request", string(data)}, " ")
}
