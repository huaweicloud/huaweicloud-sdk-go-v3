package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LtsInfo LTS信息。
type LtsInfo struct {
	Job *LtsInfoJob `json:"job,omitempty"`
}

func (o LtsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsInfo struct{}"
	}

	return strings.Join([]string{"LtsInfo", string(data)}, " ")
}
