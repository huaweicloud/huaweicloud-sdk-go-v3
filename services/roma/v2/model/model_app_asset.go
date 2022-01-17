package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppAsset struct {
}

func (o AppAsset) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppAsset struct{}"
	}

	return strings.Join([]string{"AppAsset", string(data)}, " ")
}
