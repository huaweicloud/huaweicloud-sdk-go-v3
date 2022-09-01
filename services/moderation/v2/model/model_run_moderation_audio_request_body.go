package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 请求体
type RunModerationAudioRequestBody struct {

	// 与url二选一  语音文件Base64编码字符串。要求base64编码后大小不超过4M。语音时长不超过1分钟。
	Data *string `json:"data,omitempty" xml:"data"`

	// 与data二选一  语音的URL路径，目前支持对服务授权访问华为云上OBS的URL，华为云上OBS提供的临时授权访问的URL和匿名公开授权的URL。 OBS服务的访问权限设置请参见配置OBS访问权限。 出于安全的考虑，当前服务不支持从公网上任意URL读取数据。
	Url *string `json:"url,omitempty" xml:"url"`

	Config *RunModerationAudioRequestBodyConfig `json:"config" xml:"config"`

	// 审核场景。 当前支持的场景有默认场景和用户自定义场景： ● 默认场景为：   – politics：涉政   – porn：涉黄   – ad：广告   – abuse：辱骂   – contraband：违禁品 ● 用户自定义场景为：自定义词库。 说明 自定义词库的创建和使用请参见配置自定义词库。
	Categories *[]string `json:"categories,omitempty" xml:"categories"`
}

func (o RunModerationAudioRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunModerationAudioRequestBody struct{}"
	}

	return strings.Join([]string{"RunModerationAudioRequestBody", string(data)}, " ")
}
