package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Kibana公网访问信息。
type PublicKibanaRespBody struct {

	// 带宽大小。单位：Mbit/s
	EipSize *int32 `json:"eipSize,omitempty" xml:"eipSize"`

	ElbWhiteListResp *KibanaElbWhiteListResp `json:"elbWhiteListResp,omitempty" xml:"elbWhiteListResp"`

	// kibana访问IP。
	PublicKibanaIp *string `json:"publicKibanaIp,omitempty" xml:"publicKibanaIp"`
}

func (o PublicKibanaRespBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicKibanaRespBody struct{}"
	}

	return strings.Join([]string{"PublicKibanaRespBody", string(data)}, " ")
}
