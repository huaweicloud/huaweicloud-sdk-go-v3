package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ArrayNode struct {
}

func (o ArrayNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArrayNode struct{}"
	}

	return strings.Join([]string{"ArrayNode", string(data)}, " ")
}
