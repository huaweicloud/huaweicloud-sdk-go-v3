package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCesDimsInfoRequest Request Object
type ListCesDimsInfoRequest struct {

	// 命名空间，如 SYS.LIVE，与服务上报指标时使用的namespace一致。
	Namespace string `json:"namespace"`
}

func (o ListCesDimsInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCesDimsInfoRequest struct{}"
	}

	return strings.Join([]string{"ListCesDimsInfoRequest", string(data)}, " ")
}
