package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadCertsRequest Request Object
type UploadCertsRequest struct {

	// 指定待操作的集群ID。
	ClusterId string `json:"cluster_id"`

	Body *UploadCertsRequestBody `json:"body,omitempty"`
}

func (o UploadCertsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadCertsRequest struct{}"
	}

	return strings.Join([]string{"UploadCertsRequest", string(data)}, " ")
}
