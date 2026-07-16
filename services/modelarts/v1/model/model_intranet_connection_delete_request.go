package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IntranetConnectionDeleteRequest struct {

	// 内网接入的id列表
	IntranetConnectionIds []string `json:"intranet_connection_ids"`
}

func (o IntranetConnectionDeleteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntranetConnectionDeleteRequest struct{}"
	}

	return strings.Join([]string{"IntranetConnectionDeleteRequest", string(data)}, " ")
}
