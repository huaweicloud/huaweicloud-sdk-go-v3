package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateStructConfigRequest struct {
	Body *StructConfig `json:"body,omitempty" xml:"body"`
}

func (o CreateStructConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStructConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateStructConfigRequest", string(data)}, " ")
}
