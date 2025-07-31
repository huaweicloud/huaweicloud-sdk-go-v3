package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterId 集群ID
type ClusterId struct {
}

func (o ClusterId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterId struct{}"
	}

	return strings.Join([]string{"ClusterId", string(data)}, " ")
}
