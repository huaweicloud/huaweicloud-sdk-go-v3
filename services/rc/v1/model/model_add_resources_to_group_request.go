package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddResourcesToGroupRequest Request Object
type AddResourcesToGroupRequest struct {

	// 资源组ID
	GroupId string `json:"group_id"`

	Body *AddResourcesRequestBody `json:"body,omitempty"`
}

func (o AddResourcesToGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddResourcesToGroupRequest struct{}"
	}

	return strings.Join([]string{"AddResourcesToGroupRequest", string(data)}, " ")
}
