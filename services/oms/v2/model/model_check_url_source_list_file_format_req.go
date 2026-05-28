package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckUrlSourceListFileFormatReq 检查url来源列表文件格式请求体
type CheckUrlSourceListFileFormatReq struct {

	// 存放对象列表文件的OBS桶名。 请确保与目的端桶处于同一区域，否则将导致任务创建失败。
	ObsBucket string `json:"obs_bucket"`

	// 对象列表文件或URL列表文件对象名。
	ListFileKey string `json:"list_file_key"`

	// 目的端桶的AK（最大长度100个字符）。
	Ak string `json:"ak"`

	// 目的端桶的SK（最大长度100个字符）。
	Sk string `json:"sk"`

	// 桶所处的区域。
	Region string `json:"region"`
}

func (o CheckUrlSourceListFileFormatReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckUrlSourceListFileFormatReq struct{}"
	}

	return strings.Join([]string{"CheckUrlSourceListFileFormatReq", string(data)}, " ")
}
