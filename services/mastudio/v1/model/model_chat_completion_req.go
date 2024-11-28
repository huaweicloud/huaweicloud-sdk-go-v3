package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChatCompletionReq struct {

	// 多轮对话问答对
	Messages []Message `json:"messages"`

	// 用于代表客户的唯一标识符，最小长度：1，最大长度64
	User *string `json:"user,omitempty"`

	// 流式调用的开启开关，true为开启流式调用，如果要开启流式调用，请使用流式SDK；false为关闭流式调用。默认为关闭状态（当前API Explorer不支持流式，在API Explorer调试时请使用非流式）。
	Stream *bool `json:"stream,omitempty"`

	// 用于控制生成文本的多样性和创造力。参数的取值范围是0到1，其中0表示最低的随机性。一般来说，temperature越低，适合完成确定性的任务。temperature越高，例如0.9，适合完成创造性的任务。temperature参数可以影响语言模型输出的质量和多样性，但也不是唯一的因素。还有其他一些参数，如top_p参数也可以用来调整语言模型的行为和偏好，但不建议同时更改这两个参数。
	Temperature *float32 `json:"temperature,omitempty"`

	// 一种替代温度采样的方法，称为nucleus sampling，取值范围：(0, 1]，其中模型考虑具有top_p 概率质量的标记的结果。因此 0.1 意味着只考虑构成前 10% 概率质量的标记。我们通常建议更改此值或温度，但不要同时更改两者。通常建议更改top_p或temperature来调整生成文本的倾向性，但不要同时更改这两个参数。
	TopP *float32 `json:"top_p,omitempty"`

	// 用于控制聊天回复的长度和质量。一般来说，较大的max_tokens值可以生成较长和较完整的回复，但也可能增加生成无关或重复内容的风险。较小的max_tokens值可以生成较短和较简洁的回复，但也可能导致生成不完整或不连贯的内容。因此，需要根据不同的场景和需求来选择合适的max_tokens值。最小值：1，最大值：根据模型不同最大值不同。
	MaxTokens *int32 `json:"max_tokens,omitempty"`

	// 表示对每个问题生成多少条答案。n参数的默认值是1，表示只生成一个答案。如果想要生成多条答案，可以设置n参数为一个大于1的整数，例如n=2。这样，API会返回一个包含2个答案的数组。流式调用时，n只能取1。最小值：1，最大值：2，默认值：1
	N *int32 `json:"n,omitempty"`

	// 用于控制生成文本中的重复程度。正值会根据它们到目前为止在文本中的现有频率来惩罚新tokens，从而降低模型逐字重复同一行的可能性。  presence_penalty 参数可以用来提高生成文本的多样性和创造性，避免生成单调或重复的内容。最小值：-2，最大值：2
	PresencePenalty *float32 `json:"presence_penalty,omitempty"`

	// 用于调整模型对频繁出现的Token的处理方式。即如果一个Token在训练集中出现的频率较高，那么模型在生成这个Token时会受到一定的惩罚。当frequency_penalty的值为正数时，模型会更倾向于生成出现频率较低的Token，即模型会更倾向于使用不常见的词汇。最小值：-2，最大值：2
	FrequencyPenalty *float32 `json:"frequency_penalty,omitempty"`
}

func (o ChatCompletionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChatCompletionReq struct{}"
	}

	return strings.Join([]string{"ChatCompletionReq", string(data)}, " ")
}
