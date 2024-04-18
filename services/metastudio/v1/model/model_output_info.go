package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OutputInfo 输出信息。
type OutputInfo struct {

	// 面部表情输入地址。
	FaceAddr *string `json:"face_addr,omitempty"`

	// 身体动作输入地址。
	BodyAddr *string `json:"body_addr,omitempty"`

	// 音频输入地址。
	AudioAddr *string `json:"audio_addr,omitempty"`

	// 会话ID。
	SessionId *int32 `json:"session_id,omitempty"`

	// 输出数据的格式版本，如请求中无此参数，则输出数据格式为1.0，可选值有： 1.0: 对应的输出为：         动作数据：75个骨骼旋转值         表情数据：52ARkit表情及参数 2.0: 对应的输出为：         动作数据：55个骨骼旋转值+骨骼3D坐标         表情数据：178个控制器的数据
	OutputDataVersion *string `json:"output_data_version,omitempty"`
}

func (o OutputInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputInfo struct{}"
	}

	return strings.Join([]string{"OutputInfo", string(data)}, " ")
}
