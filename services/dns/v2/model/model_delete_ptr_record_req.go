package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeletePtrRecordReq struct {

	// 待删除PTR ID。
	FloatingipIds *[]string `json:"floatingip_ids,omitempty"`
}

func (o DeletePtrRecordReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePtrRecordReq struct{}"
	}

	return strings.Join([]string{"DeletePtrRecordReq", string(data)}, " ")
}
