package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisassociateCertsV2Response Response Object
type BatchDisassociateCertsV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDisassociateCertsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisassociateCertsV2Response struct{}"
	}

	return strings.Join([]string{"BatchDisassociateCertsV2Response", string(data)}, " ")
}
