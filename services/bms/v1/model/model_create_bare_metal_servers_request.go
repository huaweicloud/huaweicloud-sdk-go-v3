package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBareMetalServersRequest Request Object
type CreateBareMetalServersRequest struct {
	Body *CreateBaremetalServersBody `json:"body,omitempty"`
}

func (o CreateBareMetalServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBareMetalServersRequest struct{}"
	}

	return strings.Join([]string{"CreateBareMetalServersRequest", string(data)}, " ")
}
