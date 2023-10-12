package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileRedirectionOptions 文件重定向控制的选项。
type FileRedirectionOptions struct {

	// 是否开启流控开关。取值为： false：表示关闭。 true：表示开启。
	FluidControlSwitchEnable *bool `json:"fluid_control_switch_enable,omitempty"`

	FluidControlOptions *FileRedirectionOptionsFluidControlOptions `json:"fluid_control_options,omitempty"`

	// 是否开启压缩开关。取值为： false：表示关闭。 true：表示开启。
	CompressionSwitchEnable *bool `json:"compression_switch_enable,omitempty"`

	CompressionSwitchOptions *FileRedirectionOptionsCompressionSwitchOptions `json:"compression_switch_options,omitempty"`

	// 是否开启Linux支持设置文件大小。取值为： false：表示关闭。 true：表示开启。
	LinuxFileSizeSupportedEnable *bool `json:"linux_file_size_supported_enable,omitempty"`

	LinuxFileSizeSupportedOptions *FileRedirectionOptionsLinuxFileSizeSupportedOptions `json:"linux_file_size_supported_options,omitempty"`

	// 是否开启Linux根目录挂载开关。取值为： false：表示关闭。 true：表示开启。
	LinuxRootMountSwitchEnable *bool `json:"linux_root_mount_switch_enable,omitempty"`

	// Linux根目录挂载路径。默认：\"\\home\"。
	LinuxRootDirList *string `json:"linux_root_dir_list,omitempty"`

	// Linux文件系统挂载路径。默认：\"\\media|\\Volumes|\\swdb\\mnt|\\home|\\storage|\\tmp|\\run\\media\"。
	LinuxFileMountPath *string `json:"linux_file_mount_path,omitempty"`

	// Linux固定驱动器文件系统格式。
	LinuxFixedDriveFileSystemFormat *string `json:"linux_fixed_drive_file_system_format,omitempty"`

	// Linux可移动驱动器文件系统格式。默认：\"vfat|ntfs|msdos|fuseblk|sdcardfs|exfat|fuse.fdredir\"。
	LinuxRemovableDriveFileSystemFormat *string `json:"linux_removable_drive_file_system_format,omitempty"`

	// Linux光盘驱动器文件系统格式。默认：\"cd9660|iso9660|udf\"。
	LinuxCdromDriveFileSystemFormat *string `json:"linux_cdrom_drive_file_system_format,omitempty"`

	// Linux网络驱动器文件系统格式。默认：\"smbfs|afpfs|cifs\"。
	LinuxNetworkDriveFileSystemFormat *string `json:"linux_network_drive_file_system_format,omitempty"`

	// 路径分隔符。默认：\"|\"。
	PathSeparator *string `json:"path_separator,omitempty"`

	// 是否开启固定驱动器（如: 本地磁盘）。取值为： false：表示关闭。 true：表示开启。
	FixedDriveEnable *bool `json:"fixed_drive_enable,omitempty"`

	// 是否开启可移除驱动器（如: U盘）。取值为： false：表示关闭。 true：表示开启。
	RemovableDriveEnable *bool `json:"removable_drive_enable,omitempty"`

	// 是否开启光盘驱动器。取值为： false：表示关闭。 true：表示开启。
	CdRomDriveEnable *bool `json:"cd_rom_drive_enable,omitempty"`

	// 是否开启网络驱动器。取值为： false：表示关闭。 true：表示开启。
	NetworkDriveEnable *bool `json:"network_drive_enable,omitempty"`
}

func (o FileRedirectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileRedirectionOptions struct{}"
	}

	return strings.Join([]string{"FileRedirectionOptions", string(data)}, " ")
}
