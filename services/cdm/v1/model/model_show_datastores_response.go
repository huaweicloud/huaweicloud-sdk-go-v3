package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatastoresResponse Response Object
type ShowDatastoresResponse struct {

	// 数据库列表。
	Datastores     *[]CdmClusterDatastore `json:"datastores,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowDatastoresResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatastoresResponse struct{}"
	}

	return strings.Join([]string{"ShowDatastoresResponse", string(data)}, " ")
}
