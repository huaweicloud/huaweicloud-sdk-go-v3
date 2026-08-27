package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserGroupsNewRequest Request Object
type ExportUserGroupsNewRequest struct {
	Body *ExportUserGroupsNewReq `json:"body,omitempty"`
}

func (o ExportUserGroupsNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserGroupsNewRequest struct{}"
	}

	return strings.Join([]string{"ExportUserGroupsNewRequest", string(data)}, " ")
}
