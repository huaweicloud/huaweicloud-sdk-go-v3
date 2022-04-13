package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddNodeRequest struct {
	// 集群 ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。

	ClusterId string `json:"cluster_id"`

	Body *AddNodeList `json:"body,omitempty"`
}

func (o AddNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddNodeRequest struct{}"
	}

	return strings.Join([]string{"AddNodeRequest", string(data)}, " ")
}
