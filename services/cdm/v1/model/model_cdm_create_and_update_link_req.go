package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CdmCreateAndUpdateLinkReq struct {
	// 连接列表，请参见links数据结构说明

	Links []Links `json:"links"`
}

func (o CdmCreateAndUpdateLinkReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmCreateAndUpdateLinkReq struct{}"
	}

	return strings.Join([]string{"CdmCreateAndUpdateLinkReq", string(data)}, " ")
}
