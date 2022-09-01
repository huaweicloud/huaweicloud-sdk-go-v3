package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteIterationsV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id" xml:"project_id"`

	Body *BatchDeleteIterationsV4RequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchDeleteIterationsV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIterationsV4Request struct{}"
	}

	return strings.Join([]string{"BatchDeleteIterationsV4Request", string(data)}, " ")
}
