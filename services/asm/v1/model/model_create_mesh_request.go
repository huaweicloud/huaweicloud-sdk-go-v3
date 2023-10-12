package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMeshRequest Request Object
type CreateMeshRequest struct {

	// 网格所属ProjectID
	XApplyProjectID *string `json:"X-Apply-ProjectID,omitempty"`

	Body *Mesh `json:"body,omitempty"`
}

func (o CreateMeshRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMeshRequest struct{}"
	}

	return strings.Join([]string{"CreateMeshRequest", string(data)}, " ")
}
