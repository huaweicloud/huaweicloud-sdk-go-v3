package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultipartFile struct {
}

func (o MultipartFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultipartFile struct{}"
	}

	return strings.Join([]string{"MultipartFile", string(data)}, " ")
}
