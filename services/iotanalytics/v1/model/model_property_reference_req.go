package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PropertyReferenceReq 属性引用
type PropertyReferenceReq struct {

	// 引用的资产ID，只有single型参数才能填写，可填写null置空
	AssetId *string `json:"asset_id,omitempty"`
}

func (o PropertyReferenceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropertyReferenceReq struct{}"
	}

	return strings.Join([]string{"PropertyReferenceReq", string(data)}, " ")
}
