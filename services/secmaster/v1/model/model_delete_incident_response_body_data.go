package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIncidentResponseBodyData 批量删除事件返回对象
type DeleteIncidentResponseBodyData struct {

	// 失败id
	ErrorIds *[]string `json:"error_ids,omitempty"`

	// 成功id
	SuccessIds *[]string `json:"success_ids,omitempty"`
}

func (o DeleteIncidentResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIncidentResponseBodyData struct{}"
	}

	return strings.Join([]string{"DeleteIncidentResponseBodyData", string(data)}, " ")
}
