package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceReqDetail 服务详情
type ServiceReqDetail struct {
	MetaData *SvcMetadata `json:"meta_data"`

	Spec *SvcSpec `json:"spec"`
}

func (o ServiceReqDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceReqDetail struct{}"
	}

	return strings.Join([]string{"ServiceReqDetail", string(data)}, " ")
}
