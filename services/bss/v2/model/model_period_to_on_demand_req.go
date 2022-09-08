package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeriodToOnDemandReq struct {

	// 设置或取消包年/包月资源到期转按需的操作。 SET_UP：设置CANCEL：取消
	Operation string `json:"operation"`

	// 资源ID。 您可以调用“查询客户包年/包月资源列表”接口获取资源ID。 此处只支持设置主资源ID，最多可设置100个资源ID。设置后，主资源及其对应的从资源将一起转为按需资源。 请根据“查询客户包年/包月资源列表”接口响应参数中的“is_main_resource”参数来标识资源是否为主资源。
	ResourceIds []string `json:"resource_ids"`
}

func (o PeriodToOnDemandReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodToOnDemandReq struct{}"
	}

	return strings.Join([]string{"PeriodToOnDemandReq", string(data)}, " ")
}
