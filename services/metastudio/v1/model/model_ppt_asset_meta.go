package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PptAssetMeta PPT资产元数据信息。
type PptAssetMeta struct {

	// **参数解释**： PPT是否需要自动解析。 **约束限制**： 部分过于复杂的PPT或压缩比过高的PPT可能无法解析。 超过50页PPT仅转换50页 转换的图片无法保证完全还原，需要自行确认。 **取值范围**： * true: 自动解析 * false: 无需解析
	AutoAnalysis *bool `json:"auto_analysis,omitempty"`

	// **参数解释**： PPT解析状态。 **约束限制**： 不涉及 **取值范围**： * INITIALIZE：初始 * WAITING：等待 * CONVERTING：解析中 * FAILED：失败 * SUCCEEDED：成功 * CANCELED：取消  **默认取值**： 不涉及
	PptAnalysisStatus *PptAssetMetaPptAnalysisStatus `json:"ppt_analysis_status,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`

	// **参数解释**： PPT页面总数。 **约束限制**： 不涉及
	PageCount *int32 `json:"page_count,omitempty"`

	// PPT页面图片。
	Pages *[]PptPageInfo `json:"pages,omitempty"`
}

func (o PptAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PptAssetMeta struct{}"
	}

	return strings.Join([]string{"PptAssetMeta", string(data)}, " ")
}

type PptAssetMetaPptAnalysisStatus struct {
	value string
}

type PptAssetMetaPptAnalysisStatusEnum struct {
	INITIALIZE PptAssetMetaPptAnalysisStatus
	WAITING    PptAssetMetaPptAnalysisStatus
	CONVERTING PptAssetMetaPptAnalysisStatus
	FAILED     PptAssetMetaPptAnalysisStatus
	SUCCEEDED  PptAssetMetaPptAnalysisStatus
	CANCELED   PptAssetMetaPptAnalysisStatus
}

func GetPptAssetMetaPptAnalysisStatusEnum() PptAssetMetaPptAnalysisStatusEnum {
	return PptAssetMetaPptAnalysisStatusEnum{
		INITIALIZE: PptAssetMetaPptAnalysisStatus{
			value: "INITIALIZE",
		},
		WAITING: PptAssetMetaPptAnalysisStatus{
			value: "WAITING",
		},
		CONVERTING: PptAssetMetaPptAnalysisStatus{
			value: "CONVERTING",
		},
		FAILED: PptAssetMetaPptAnalysisStatus{
			value: "FAILED",
		},
		SUCCEEDED: PptAssetMetaPptAnalysisStatus{
			value: "SUCCEEDED",
		},
		CANCELED: PptAssetMetaPptAnalysisStatus{
			value: "CANCELED",
		},
	}
}

func (c PptAssetMetaPptAnalysisStatus) Value() string {
	return c.value
}

func (c PptAssetMetaPptAnalysisStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PptAssetMetaPptAnalysisStatus) UnmarshalJSON(b []byte) error {
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
