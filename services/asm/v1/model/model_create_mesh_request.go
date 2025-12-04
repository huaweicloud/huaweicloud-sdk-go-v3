package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMeshRequest Request Object
type CreateMeshRequest struct {
	Body *Mesh `json:"body,omitempty"`
}

func (o CreateMeshRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMeshRequest struct{}"
	}

	return strings.Join([]string{"CreateMeshRequest", string(data)}, " ")
}
