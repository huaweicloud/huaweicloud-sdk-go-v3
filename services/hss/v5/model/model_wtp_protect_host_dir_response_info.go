package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type WtpProtectHostDirResponseInfo struct {

	// 防护目录
	ProtectDir *string `json:"protect_dir,omitempty"`

	// 排除子目录
	ExcludeChildDir *string `json:"exclude_child_dir,omitempty"`

	// 排除文件路径
	ExcludeFilePath *string `json:"exclude_file_path,omitempty"`

	// 排除文件路径
	ExclueFilePath *string `json:"exclue_file_path,omitempty"`

	// 本地备份路径
	LocalBackupDir *string `json:"local_backup_dir,omitempty"`

	// 防护状态（closed-未开启，opened-防护中, opening-开启中, closing-关闭中, open_failed-防护失败)
	ProtectStatus *WtpProtectHostDirResponseInfoProtectStatus `json:"protect_status,omitempty"`

	// 失败原因
	Error *string `json:"error,omitempty"`
}

func (o WtpProtectHostDirResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WtpProtectHostDirResponseInfo struct{}"
	}

	return strings.Join([]string{"WtpProtectHostDirResponseInfo", string(data)}, " ")
}

type WtpProtectHostDirResponseInfoProtectStatus struct {
	value string
}

type WtpProtectHostDirResponseInfoProtectStatusEnum struct {
	CLOSED      WtpProtectHostDirResponseInfoProtectStatus
	OPENED      WtpProtectHostDirResponseInfoProtectStatus
	OPENING     WtpProtectHostDirResponseInfoProtectStatus
	CLOSING     WtpProtectHostDirResponseInfoProtectStatus
	OPEN_FAILED WtpProtectHostDirResponseInfoProtectStatus
}

func GetWtpProtectHostDirResponseInfoProtectStatusEnum() WtpProtectHostDirResponseInfoProtectStatusEnum {
	return WtpProtectHostDirResponseInfoProtectStatusEnum{
		CLOSED: WtpProtectHostDirResponseInfoProtectStatus{
			value: "closed",
		},
		OPENED: WtpProtectHostDirResponseInfoProtectStatus{
			value: "opened",
		},
		OPENING: WtpProtectHostDirResponseInfoProtectStatus{
			value: "opening",
		},
		CLOSING: WtpProtectHostDirResponseInfoProtectStatus{
			value: "closing",
		},
		OPEN_FAILED: WtpProtectHostDirResponseInfoProtectStatus{
			value: "open_failed",
		},
	}
}

func (c WtpProtectHostDirResponseInfoProtectStatus) Value() string {
	return c.value
}

func (c WtpProtectHostDirResponseInfoProtectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WtpProtectHostDirResponseInfoProtectStatus) UnmarshalJSON(b []byte) error {
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
