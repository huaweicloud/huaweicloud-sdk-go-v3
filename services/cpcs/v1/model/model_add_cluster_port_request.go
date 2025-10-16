package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddClusterPortRequest Request Object
type AddClusterPortRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *CreateElbClusterPortRequestBody `json:"body,omitempty"`
}

func (o AddClusterPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddClusterPortRequest struct{}"
	}

	return strings.Join([]string{"AddClusterPortRequest", string(data)}, " ")
}
