package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAlterKibanaRequest struct {
	// 指定待更改的集群ID。

	ClusterId string `json:"cluster_id"`

	Body *UpdatePublicKibanaBandwidthReq `json:"body,omitempty"`
}

func (o UpdateAlterKibanaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlterKibanaRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlterKibanaRequest", string(data)}, " ")
}
