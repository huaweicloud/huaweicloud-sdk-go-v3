package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateAppsResponse Response Object
type DisassociateAppsResponse struct {

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 应用id列表
	AppId          *[]string `json:"app_id,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DisassociateAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateAppsResponse struct{}"
	}

	return strings.Join([]string{"DisassociateAppsResponse", string(data)}, " ")
}
