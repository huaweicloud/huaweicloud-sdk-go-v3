package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlameLineTreeResponse Response Object
type ShowFlameLineTreeResponse struct {

	// 火焰图的数据，是个二维数组 data[0][0]: self time，方法自己消耗的cpu毫秒时间，不包括方法内部调用其他方法的时间 data[0][1]: total time, 方法消耗的cpu毫秒时间，包括方法内部调用其他方法的时间 data[0][2]: 方法的index,对应methods中的数组下标 data[0][3]: 行号 data[0][4]: 方法的子节点，就是方法中调用的其他方法
	Data *[]string `json:"data,omitempty"`

	// 调用栈上的方法信息，是个二维数组 method[0][0]: 方法的唯一id method[0][1]: 方法的package包名 method[0][2]: 方法的class name 类名 method[0][3]: 方法名 method[0][4]: 方法的参数列表 method[0][5]: 方法是否为用户的方法 method[0][6]: 方法是否为native方法
	Methods        *[]string `json:"methods,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowFlameLineTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlameLineTreeResponse struct{}"
	}

	return strings.Join([]string{"ShowFlameLineTreeResponse", string(data)}, " ")
}
