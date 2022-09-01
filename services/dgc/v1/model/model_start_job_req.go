package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartJobReq struct {
	JobParams *[]JobParam `json:"jobParams,omitempty" xml:"jobParams"`
}

func (o StartJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartJobReq struct{}"
	}

	return strings.Join([]string{"StartJobReq", string(data)}, " ")
}
