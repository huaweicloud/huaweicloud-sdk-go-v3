package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddVolumesReq struct {

	// 订单ID，包周期桌面添加磁盘时使用。
	OrderId *string `json:"order_id,omitempty"`

	// 企业项目ID，默认\"0\"
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 待新增的磁盘信息，每个桌面的数据盘数量不超过10个。
	Volumes []Volume `json:"volumes"`
}

func (o AddVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVolumesReq struct{}"
	}

	return strings.Join([]string{"AddVolumesReq", string(data)}, " ")
}
