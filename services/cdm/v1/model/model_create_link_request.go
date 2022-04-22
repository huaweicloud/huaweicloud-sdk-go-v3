package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateLinkRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 为“true”时，此API仅校验参数是否正确，不创建连接
	Validate *string `json:"validate,omitempty"`

	Body *CdmCreateAndUpdateLinkReq `json:"body,omitempty"`
}

func (o CreateLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLinkRequest struct{}"
	}

	return strings.Join([]string{"CreateLinkRequest", string(data)}, " ")
}
