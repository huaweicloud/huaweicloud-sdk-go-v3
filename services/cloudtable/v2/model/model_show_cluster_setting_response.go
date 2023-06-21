package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowClusterSettingResponse struct {

	// 集群参数生效状态：0、未更改 1、未应用 2、应用中 3、已应用 4、应用失败
	ParmStatus *int32 `json:"parm_status,omitempty"`

	// 参数列表
	ParameterInfo  *[]ParameterInfo `json:"parameter_info,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowClusterSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterSettingResponse", string(data)}, " ")
}
