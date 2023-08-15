package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestorePtrRecordRequest Request Object
type RestorePtrRecordRequest struct {

	// 域名所属的区域。
	Region string `json:"region"`

	// 待删除PTR ID。
	FloatingipId string `json:"floatingip_id"`

	Body *RestorePtrReq `json:"body,omitempty"`
}

func (o RestorePtrRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestorePtrRecordRequest struct{}"
	}

	return strings.Join([]string{"RestorePtrRecordRequest", string(data)}, " ")
}
