package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDatabaseAuthorityRequest Request Object
type UpdateDatabaseAuthorityRequest struct {

	// cluster_id
	ClusterId string `json:"cluster_id"`

	Body *DatabasePermissionReq `json:"body,omitempty"`
}

func (o UpdateDatabaseAuthorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseAuthorityRequest struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseAuthorityRequest", string(data)}, " ")
}
