package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDatabaseUserInfoRequest Request Object
type UpdateDatabaseUserInfoRequest struct {

	// cluster_id
	ClusterId string `json:"cluster_id"`

	// name
	Name string `json:"name"`

	Body *DatabaseUserInfoReq `json:"body,omitempty"`
}

func (o UpdateDatabaseUserInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseUserInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseUserInfoRequest", string(data)}, " ")
}
