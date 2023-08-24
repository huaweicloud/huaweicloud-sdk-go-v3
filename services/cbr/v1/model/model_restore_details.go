package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreDetails
type RestoreDetails struct {

	// 目的路径
	DestinationPath string `json:"destination_path"`
}

func (o RestoreDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreDetails struct{}"
	}

	return strings.Join([]string{"RestoreDetails", string(data)}, " ")
}
