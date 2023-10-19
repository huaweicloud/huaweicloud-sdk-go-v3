package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSpec struct {

	// 商品ID。由系统自动生成，如OFFI8XXXXXXXXXXXXXXXX4。
	ProductId *string `json:"product_id,omitempty"`

	// 规格编码。由系统自动生成，例如lakeformation.unit.basic.qps。
	SpecCode string `json:"spec_code"`

	// 步数。QPS为每秒最大请求数步长，最小为5，步长为1。
	StrideNum int32 `json:"stride_num"`
}

func (o CreateSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSpec struct{}"
	}

	return strings.Join([]string{"CreateSpec", string(data)}, " ")
}
