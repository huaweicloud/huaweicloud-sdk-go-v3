package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeExclusiveLitesReq 专享主机变更桌面桌面路数请求。
type ResizeExclusiveLitesReq struct {

	// 企业项目ID，默认\"0。\"
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 产品套餐ID。
	ProductId string `json:"product_id"`

	// 扩容后的桌面个数，单位为个/次。
	NewQuantity int32 `json:"new_quantity"`
}

func (o ResizeExclusiveLitesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeExclusiveLitesReq struct{}"
	}

	return strings.Join([]string{"ResizeExclusiveLitesReq", string(data)}, " ")
}
