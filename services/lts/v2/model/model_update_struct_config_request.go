package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateStructConfigRequest struct {
	Body *StructConfig `json:"body,omitempty" xml:"body"`
}

func (o UpdateStructConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStructConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateStructConfigRequest", string(data)}, " ")
}
