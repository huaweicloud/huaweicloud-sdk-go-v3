package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMeshesRequest Request Object
type ListMeshesRequest struct {

	// 网格所属ProjectID
	XApplyProjectID *string `json:"X-Apply-ProjectID,omitempty"`
}

func (o ListMeshesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMeshesRequest struct{}"
	}

	return strings.Join([]string{"ListMeshesRequest", string(data)}, " ")
}
