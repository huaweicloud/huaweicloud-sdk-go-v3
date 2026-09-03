package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SignProtocolNewRequest Request Object
type SignProtocolNewRequest struct {
}

func (o SignProtocolNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignProtocolNewRequest struct{}"
	}

	return strings.Join([]string{"SignProtocolNewRequest", string(data)}, " ")
}
