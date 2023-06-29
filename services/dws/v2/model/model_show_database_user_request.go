package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatabaseUserRequest Request Object
type ShowDatabaseUserRequest struct {

	// cluster_id
	ClusterId string `json:"cluster_id"`

	// name
	Name string `json:"name"`
}

func (o ShowDatabaseUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"ShowDatabaseUserRequest", string(data)}, " ")
}
