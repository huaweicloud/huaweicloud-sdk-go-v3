package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PredictUrlResponse **参数解释：** 在线服务访问地址。
type PredictUrlResponse struct {

	// **参数解释：** 访问方式。 **取值范围：** - CONSOLE：通过控制台界面访问。 - PUBLIC：通过公网访问。 - INTERNAL：通过内网访问。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 推理请求的访问地址，仅当type为REAL_TIME时，且服务部署完成后才会确保该字段有值。
	Urls *[]string `json:"urls,omitempty"`
}

func (o PredictUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PredictUrlResponse struct{}"
	}

	return strings.Join([]string{"PredictUrlResponse", string(data)}, " ")
}
