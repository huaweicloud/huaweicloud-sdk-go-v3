package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndividualParam 个性化动态参数。
type IndividualParam struct {

	// 个性化动态参数号码列表，最多支持5000个号码。  > 长度指的是单个号码的长度。
	Mobiles []string `json:"mobiles"`

	// 个性化动态参数数组。 参数顺序按照模板创建时参数占位符的顺序传入，例如创建模板时设置动参有#p_1#、#p_2#、#p_3#，则传入的参数数组顺序第一个元素为#p_1#，第二个元素是#p_2#，第三个元素为#p_3#。
	DyncParams []IndividualContentParam `json:"dync_params"`
}

func (o IndividualParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndividualParam struct{}"
	}

	return strings.Join([]string{"IndividualParam", string(data)}, " ")
}
