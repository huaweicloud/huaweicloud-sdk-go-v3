package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// KnowledgeBaseRetrievalReq 知识检索请求体。
type KnowledgeBaseRetrievalReq struct {

	// **参数解释**： 知识库ID列表。  **约束限制**： 最多可同时检索3个知识库。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	KnowledgeBaseIds []string `json:"knowledge_base_ids"`

	// **参数解释**： 用户输入的问题或关键词。  **约束限制**： 不涉及。  **取值范围**： 长度 1 至 4096 字符。  **默认取值**： 不涉及。
	Query string `json:"query"`

	// **参数解释**： 检索策略模式。  **约束限制**： 不涉及。  **取值范围**： - doc：语义检索。 - keyword：关键词检索。 - mix：混合检索。 - faq：FAQ检索。  **默认取值**： doc。
	SearchMode *KnowledgeBaseRetrievalReqSearchMode `json:"search_mode,omitempty"`

	// **参数解释**： 每个知识库最多返回的检索结果数量。  **约束限制**： 若传入小数，系统会默认截断小数部分。  **取值范围**： 1 至 100（含）。  **默认取值**： 10。
	TopK *int32 `json:"top_k,omitempty"`

	// **参数解释**： 检索结果的最低相关度得分，低于此值的片段将被过滤。  **约束限制**： 不涉及。  **取值范围**： [0.0, 1.0]，包含两端。  **默认取值**： 0.5。
	SimilarityThreshold *float32 `json:"similarity_threshold,omitempty"`

	// **参数解释**： 知识检索结果切片中，对图片标签进行处理和保留的具体方式。  **约束限制**： 该功能要求被检索的知识库本身支持返回图片信息。  **取值范围**： - RETAIN_IMAGE_ID：保留图片ID，格式：{KI|imageId}。 - RETAIN_PLACEHOLDER：保留占位符，格式：{KI|N}，N为序号。 - REMOVE_IMAGE：移除图片（即替换为空字符串）。  **默认取值**： REMOVE_IMAGE。
	ImageMaskPolicy *KnowledgeBaseRetrievalReqImageMaskPolicy `json:"image_mask_policy,omitempty"`
}

func (o KnowledgeBaseRetrievalReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KnowledgeBaseRetrievalReq struct{}"
	}

	return strings.Join([]string{"KnowledgeBaseRetrievalReq", string(data)}, " ")
}

type KnowledgeBaseRetrievalReqSearchMode struct {
	value string
}

type KnowledgeBaseRetrievalReqSearchModeEnum struct {
	DOC     KnowledgeBaseRetrievalReqSearchMode
	KEYWORD KnowledgeBaseRetrievalReqSearchMode
	MIX     KnowledgeBaseRetrievalReqSearchMode
	FAQ     KnowledgeBaseRetrievalReqSearchMode
}

func GetKnowledgeBaseRetrievalReqSearchModeEnum() KnowledgeBaseRetrievalReqSearchModeEnum {
	return KnowledgeBaseRetrievalReqSearchModeEnum{
		DOC: KnowledgeBaseRetrievalReqSearchMode{
			value: "doc",
		},
		KEYWORD: KnowledgeBaseRetrievalReqSearchMode{
			value: "keyword",
		},
		MIX: KnowledgeBaseRetrievalReqSearchMode{
			value: "mix",
		},
		FAQ: KnowledgeBaseRetrievalReqSearchMode{
			value: "faq",
		},
	}
}

func (c KnowledgeBaseRetrievalReqSearchMode) Value() string {
	return c.value
}

func (c KnowledgeBaseRetrievalReqSearchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KnowledgeBaseRetrievalReqSearchMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type KnowledgeBaseRetrievalReqImageMaskPolicy struct {
	value string
}

type KnowledgeBaseRetrievalReqImageMaskPolicyEnum struct {
	RETAIN_IMAGE_ID    KnowledgeBaseRetrievalReqImageMaskPolicy
	RETAIN_PLACEHOLDER KnowledgeBaseRetrievalReqImageMaskPolicy
	REMOVE_IMAGE       KnowledgeBaseRetrievalReqImageMaskPolicy
}

func GetKnowledgeBaseRetrievalReqImageMaskPolicyEnum() KnowledgeBaseRetrievalReqImageMaskPolicyEnum {
	return KnowledgeBaseRetrievalReqImageMaskPolicyEnum{
		RETAIN_IMAGE_ID: KnowledgeBaseRetrievalReqImageMaskPolicy{
			value: "RETAIN_IMAGE_ID",
		},
		RETAIN_PLACEHOLDER: KnowledgeBaseRetrievalReqImageMaskPolicy{
			value: "RETAIN_PLACEHOLDER",
		},
		REMOVE_IMAGE: KnowledgeBaseRetrievalReqImageMaskPolicy{
			value: "REMOVE_IMAGE",
		},
	}
}

func (c KnowledgeBaseRetrievalReqImageMaskPolicy) Value() string {
	return c.value
}

func (c KnowledgeBaseRetrievalReqImageMaskPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KnowledgeBaseRetrievalReqImageMaskPolicy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
