package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsageInfosRequest Request Object
type ListUsageInfosRequest struct {

	// 项目ID
	ProjectUuid string `json:"project_uuid"`
}

func (o ListUsageInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsageInfosRequest struct{}"
	}

	return strings.Join([]string{"ListUsageInfosRequest", string(data)}, " ")
}
