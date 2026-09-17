package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVisibleServicesRequest Request Object
type ListVisibleServicesRequest struct {

	// 项目ID
	ProjectUuid string `json:"project_uuid"`
}

func (o ListVisibleServicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVisibleServicesRequest struct{}"
	}

	return strings.Join([]string{"ListVisibleServicesRequest", string(data)}, " ")
}
