package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PreviewAgencyLogAccessReqListBody struct {

	// 预览代理列表
	PreviewAgencyList []PreviewAgencyLogAccessReqBody `json:"preview_agency_list"`
}

func (o PreviewAgencyLogAccessReqListBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviewAgencyLogAccessReqListBody struct{}"
	}

	return strings.Join([]string{"PreviewAgencyLogAccessReqListBody", string(data)}, " ")
}
