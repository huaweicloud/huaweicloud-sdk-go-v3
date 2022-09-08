package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type RestoreClusterRequestBody struct {
	Restore *Restore `json:"restore"`
}

func (o RestoreClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreClusterRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreClusterRequestBody", string(data)}, " ")
}
