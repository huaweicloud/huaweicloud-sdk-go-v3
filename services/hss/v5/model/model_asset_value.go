package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssetValue 资产重要性，包含如下3种   - important ：重要资产   - common ：一般资产   - test ：测试资产
type AssetValue struct {
}

func (o AssetValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetValue struct{}"
	}

	return strings.Join([]string{"AssetValue", string(data)}, " ")
}
