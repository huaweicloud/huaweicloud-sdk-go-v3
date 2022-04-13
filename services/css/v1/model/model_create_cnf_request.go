package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateCnfRequest struct {
	// 指定待创建的集群ID。

	ClusterId string `json:"cluster_id"`

	Body *CreateCnfReq `json:"body,omitempty"`
}

func (o CreateCnfRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCnfRequest struct{}"
	}

	return strings.Join([]string{"CreateCnfRequest", string(data)}, " ")
}
