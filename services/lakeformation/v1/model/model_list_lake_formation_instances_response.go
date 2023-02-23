package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListLakeFormationInstancesResponse struct {

	// LakeFormation实例列表
	Instances *[]LakeFormationInstance `json:"instances,omitempty"`

	// LakeFormation实例总数
	Total *int32 `json:"total,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListLakeFormationInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLakeFormationInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListLakeFormationInstancesResponse", string(data)}, " ")
}
