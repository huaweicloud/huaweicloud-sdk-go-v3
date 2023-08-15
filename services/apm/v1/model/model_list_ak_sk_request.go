package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAkSkRequest Request Object
type ListAkSkRequest struct {
}

func (o ListAkSkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAkSkRequest struct{}"
	}

	return strings.Join([]string{"ListAkSkRequest", string(data)}, " ")
}
