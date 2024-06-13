package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AssetFileInfo 文件信息。
type AssetFileInfo struct {

	// 文件ID。
	FileId *string `json:"file_id,omitempty"`

	// 文件名创建文件时候不区分大小写，最大长度256，最小长度1。
	FileName *string `json:"file_name,omitempty"`

	// 文件内容MD5值，固定24位。
	FileMd5 *string `json:"file_md5,omitempty"`

	// 文件总的大小，最小1，最大5368709120。
	FileSize *int64 `json:"file_size,omitempty"`

	// 文件类型（默认提取文件后缀）。
	FileType *string `json:"file_type,omitempty"`

	// 文件在资产中的分类。每种资产类型包含的文件分类不同。 * MAIN：主文件 * COVER：封面文件 * PAGE：内容页图片 * SAMPLE：样例音频 * OTHER：其他文件 * WHOLE_MODEL：全模型 * USER_MODIFIED_MODEL：用户上传模型 > * 资产类型为SCENE、ANIMATION、VIDEO、IMAGE、MATERIAL时，包含MAIN、COVER和OTHER > * 资产类型为PPT时，包含MAIN、COVER、PAGE和OTHER > * 资产类型为HUMAN_MODEL时，包含MAIN、COVER和OTHER > * 资产类型为VOICE_MODEL时，包含MAIN、SAMPLE(样例音频文件)和OTHER > * 资产类型为HUMAN_MODEL_2D时，包含MAIN、COVER、SAMPLE(动作样例)和OTHER(遮罩文件) > * 资产类型为BUSINESS_CARD_TEMPLET时，包含MAIN和COVER(名片效果图)
	AssetFileCategory *string `json:"asset_file_category,omitempty"`

	// 文件下载URL，有效期为24小时。
	DownloadUrl *string `json:"download_url,omitempty"`

	// 文件状态枚举: * CREATING：文件上传中 * CREATED：文件已上传（自动审核通过） * FAILED：文件上传失败 * CANCELLED：文件上传已取消 * DELETING：文件删除中 * DELETED：文件已删除 * UPLOADED：文件已上传（尚未审核） * REVIEW：人工审核（文件已上传） * BLOCK：冻结
	State *AssetFileInfoState `json:"state,omitempty"`

	// 审核失败原因
	Reason *string `json:"reason,omitempty"`
}

func (o AssetFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetFileInfo struct{}"
	}

	return strings.Join([]string{"AssetFileInfo", string(data)}, " ")
}

type AssetFileInfoState struct {
	value string
}

type AssetFileInfoStateEnum struct {
	CREATING  AssetFileInfoState
	CREATED   AssetFileInfoState
	FAILED    AssetFileInfoState
	CANCELLED AssetFileInfoState
	DELETING  AssetFileInfoState
	DELETED   AssetFileInfoState
	UPLOADED  AssetFileInfoState
	REVIEW    AssetFileInfoState
	BLOCK     AssetFileInfoState
}

func GetAssetFileInfoStateEnum() AssetFileInfoStateEnum {
	return AssetFileInfoStateEnum{
		CREATING: AssetFileInfoState{
			value: "CREATING",
		},
		CREATED: AssetFileInfoState{
			value: "CREATED",
		},
		FAILED: AssetFileInfoState{
			value: "FAILED",
		},
		CANCELLED: AssetFileInfoState{
			value: "CANCELLED",
		},
		DELETING: AssetFileInfoState{
			value: "DELETING",
		},
		DELETED: AssetFileInfoState{
			value: "DELETED",
		},
		UPLOADED: AssetFileInfoState{
			value: "UPLOADED",
		},
		REVIEW: AssetFileInfoState{
			value: "REVIEW",
		},
		BLOCK: AssetFileInfoState{
			value: "BLOCK",
		},
	}
}

func (c AssetFileInfoState) Value() string {
	return c.value
}

func (c AssetFileInfoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssetFileInfoState) UnmarshalJSON(b []byte) error {
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
