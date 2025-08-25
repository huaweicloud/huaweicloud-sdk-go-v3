package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateAppsResponse Response Object
type AssociateAppsResponse struct {

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 应用id
	AppId          *string `json:"app_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateAppsResponse struct{}"
	}

	return strings.Join([]string{"AssociateAppsResponse", string(data)}, " ")
}
