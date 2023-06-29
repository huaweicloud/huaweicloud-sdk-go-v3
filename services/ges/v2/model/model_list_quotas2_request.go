package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotas2Request Request Object
type ListQuotas2Request struct {
}

func (o ListQuotas2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotas2Request struct{}"
	}

	return strings.Join([]string{"ListQuotas2Request", string(data)}, " ")
}
