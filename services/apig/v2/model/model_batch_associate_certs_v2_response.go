package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchAssociateCertsV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAssociateCertsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAssociateCertsV2Response struct{}"
	}

	return strings.Join([]string{"BatchAssociateCertsV2Response", string(data)}, " ")
}
