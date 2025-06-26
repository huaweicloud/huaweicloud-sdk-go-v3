package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteInstanceRequestBody struct {

	// 是否删除obs桶
	DeleteObs *bool `json:"delete_obs,omitempty"`

	// 是否删除DNS域名信息
	DeleteDns *bool `json:"delete_dns,omitempty"`
}

func (o DeleteInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteInstanceRequestBody", string(data)}, " ")
}
