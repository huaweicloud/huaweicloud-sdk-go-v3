package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VirusNum 新发现病毒数量
type VirusNum struct {
}

func (o VirusNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VirusNum struct{}"
	}

	return strings.Join([]string{"VirusNum", string(data)}, " ")
}
