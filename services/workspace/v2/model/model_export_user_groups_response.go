package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserGroupsResponse Response Object
type ExportUserGroupsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportUserGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserGroupsResponse struct{}"
	}

	return strings.Join([]string{"ExportUserGroupsResponse", string(data)}, " ")
}
