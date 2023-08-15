package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterV2Request Request Object
type CreateClusterV2Request struct {
	Body *V2CreateClusterReq `json:"body,omitempty"`
}

func (o CreateClusterV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterV2Request struct{}"
	}

	return strings.Join([]string{"CreateClusterV2Request", string(data)}, " ")
}
