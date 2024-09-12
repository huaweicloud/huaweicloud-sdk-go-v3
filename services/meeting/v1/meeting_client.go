package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/meeting/v1/model"
)

type MeetingClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewMeetingClient(hcClient *httpclient.HcHttpClient) *MeetingClient {
	return &MeetingClient{HcClient: hcClient}
}

func MeetingClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("v1.MeetingCredentials")
	return builder
}

// AddAppId 添加企业应用
//
// 企业默认管理员添加应用，添加应用后，记录返回信息，后续可通过[[执行App ID鉴权](https://support.huaweicloud.com/api-meeting/meeting_21_0311.html)](tag:hws) [[执行App ID鉴权](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0311.html)](tag:hk)获取accessToken
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) AddAppId(request *model.AddAppIdRequest) (*model.AddAppIdResponse, error) {
	requestDef := GenReqDefForAddAppId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAppIdResponse), nil
	}
}

// AddAppIdInvoker 添加企业应用
func (c *MeetingClient) AddAppIdInvoker(request *model.AddAppIdRequest) *AddAppIdInvoker {
	requestDef := GenReqDefForAddAppId()
	return &AddAppIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddCorp SP管理员创建企业
//
// 创建企业，默认管理员及分配资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) AddCorp(request *model.AddCorpRequest) (*model.AddCorpResponse, error) {
	requestDef := GenReqDefForAddCorp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddCorpResponse), nil
	}
}

// AddCorpInvoker SP管理员创建企业
func (c *MeetingClient) AddCorpInvoker(request *model.AddCorpRequest) *AddCorpInvoker {
	requestDef := GenReqDefForAddCorp()
	return &AddCorpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddCorpAdmin 添加企业管理员
//
// 企业默认管理员添加企业普通管理员。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) AddCorpAdmin(request *model.AddCorpAdminRequest) (*model.AddCorpAdminResponse, error) {
	requestDef := GenReqDefForAddCorpAdmin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddCorpAdminResponse), nil
	}
}

// AddCorpAdminInvoker 添加企业管理员
func (c *MeetingClient) AddCorpAdminInvoker(request *model.AddCorpAdminRequest) *AddCorpAdminInvoker {
	requestDef := GenReqDefForAddCorpAdmin()
	return &AddCorpAdminInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddDepartment 添加部门
//
// 企业管理员通过该接口添加部门，最多支持10级部门，每级子部门最多支持100个，默认企业最大部门数量为10000个。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) AddDepartment(request *model.AddDepartmentRequest) (*model.AddDepartmentResponse, error) {
	requestDef := GenReqDefForAddDepartment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDepartmentResponse), nil
	}
}

// AddDepartmentInvoker 添加部门
func (c *MeetingClient) AddDepartmentInvoker(request *model.AddDepartmentRequest) *AddDepartmentInvoker {
	requestDef := GenReqDefForAddDepartment()
	return &AddDepartmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddDevice 增加终端
//
// 企业管理员通过该接口添加专业会议终端。专业会议终端包括DP300/HUAWEI Bar系列/HUAWEI Board/TE系列。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) AddDevice(request *model.AddDeviceRequest) (*model.AddDeviceResponse, error) {
	requestDef := GenReqDefForAddDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDeviceResponse), nil
	}
}

// AddDeviceInvoker 增加终端
func (c *MeetingClient) AddDeviceInvoker(request *model.AddDeviceRequest) *AddDeviceInvoker {
	requestDef := GenReqDefForAddDevice()
	return &AddDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddMaterial 新增信息窗素材
//
// 新增信息窗素材（上传素材文件）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) AddMaterial(request *model.AddMaterialRequest) (*model.AddMaterialResponse, error) {
	requestDef := GenReqDefForAddMaterial()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddMaterialResponse), nil
	}
}

// AddMaterialInvoker 新增信息窗素材
func (c *MeetingClient) AddMaterialInvoker(request *model.AddMaterialRequest) *AddMaterialInvoker {
	requestDef := GenReqDefForAddMaterial()
	return &AddMaterialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddProgram 新增信息窗节目
//
// 新增信息窗节目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) AddProgram(request *model.AddProgramRequest) (*model.AddProgramResponse, error) {
	requestDef := GenReqDefForAddProgram()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddProgramResponse), nil
	}
}

// AddProgramInvoker 新增信息窗节目
func (c *MeetingClient) AddProgramInvoker(request *model.AddProgramRequest) *AddProgramInvoker {
	requestDef := GenReqDefForAddProgram()
	return &AddProgramInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddPublication 新增信息窗发布
//
// 新增信息窗发布。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) AddPublication(request *model.AddPublicationRequest) (*model.AddPublicationResponse, error) {
	requestDef := GenReqDefForAddPublication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddPublicationResponse), nil
	}
}

// AddPublicationInvoker 新增信息窗发布
func (c *MeetingClient) AddPublicationInvoker(request *model.AddPublicationRequest) *AddPublicationInvoker {
	requestDef := GenReqDefForAddPublication()
	return &AddPublicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddResource SP管理员分配企业资源
//
// 企业新增资源发放。该接口同时支持修改，带resourceId后会判断该资源是否存在，存在即修改（支持修改的参数见修改接口），否则按新增处理。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) AddResource(request *model.AddResourceRequest) (*model.AddResourceResponse, error) {
	requestDef := GenReqDefForAddResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddResourceResponse), nil
	}
}

// AddResourceInvoker SP管理员分配企业资源
func (c *MeetingClient) AddResourceInvoker(request *model.AddResourceRequest) *AddResourceInvoker {
	requestDef := GenReqDefForAddResource()
	return &AddResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddToPersonalSpace 保存会议纪要到个人云空间
//
// 用户使用手机扫码后,手机端请求服务端将当前会议纪要文件保存到个人云空间。二维码内容 ：cloudlink://cloudlink.huawei.com/h5page?action&#x3D;SAVE_MEETING_FILE&amp;key1&#x3D;value1&amp;key2&#x3D;value2 。key/value的个数可能变化，终端解析后，在发起后续请求时，将所有key/value存为map，作为入参即可。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) AddToPersonalSpace(request *model.AddToPersonalSpaceRequest) (*model.AddToPersonalSpaceResponse, error) {
	requestDef := GenReqDefForAddToPersonalSpace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddToPersonalSpaceResponse), nil
	}
}

// AddToPersonalSpaceInvoker 保存会议纪要到个人云空间
func (c *MeetingClient) AddToPersonalSpaceInvoker(request *model.AddToPersonalSpaceRequest) *AddToPersonalSpaceInvoker {
	requestDef := GenReqDefForAddToPersonalSpace()
	return &AddToPersonalSpaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddUser 添加用户
//
// 企业管理员通过该接口添加企业用户。
// &gt; 默认添加用户后，用户第一次登录华为云会议App或者Portal时需要修改密码。若需关闭第一次登录修改密码，请联系华为销售人员，并提供华为云会议企业ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) AddUser(request *model.AddUserRequest) (*model.AddUserResponse, error) {
	requestDef := GenReqDefForAddUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddUserResponse), nil
	}
}

// AddUserInvoker 添加用户
func (c *MeetingClient) AddUserInvoker(request *model.AddUserRequest) *AddUserInvoker {
	requestDef := GenReqDefForAddUser()
	return &AddUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AllowClientRecord 允许客户端录制
//
// 该接口用于设置允许/禁止与会者客户端本地录制（非云端录制）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) AllowClientRecord(request *model.AllowClientRecordRequest) (*model.AllowClientRecordResponse, error) {
	requestDef := GenReqDefForAllowClientRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AllowClientRecordResponse), nil
	}
}

// AllowClientRecordInvoker 允许客户端录制
func (c *MeetingClient) AllowClientRecordInvoker(request *model.AllowClientRecordRequest) *AllowClientRecordInvoker {
	requestDef := GenReqDefForAllowClientRecord()
	return &AllowClientRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AllowGuestUnmute 与会者自己解除静音
//
// 该接口用于设置与会者是否可以自己解除静音。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) AllowGuestUnmute(request *model.AllowGuestUnmuteRequest) (*model.AllowGuestUnmuteResponse, error) {
	requestDef := GenReqDefForAllowGuestUnmute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AllowGuestUnmuteResponse), nil
	}
}

// AllowGuestUnmuteInvoker 与会者自己解除静音
func (c *MeetingClient) AllowGuestUnmuteInvoker(request *model.AllowGuestUnmuteRequest) *AllowGuestUnmuteInvoker {
	requestDef := GenReqDefForAllowGuestUnmute()
	return &AllowGuestUnmuteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AllowWaitingParticipant 准入等候者
//
// 该接口用于允许等候室中的成员进入会议。可以允许全部成员进入会议，或者允许指定成员进入会议。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) AllowWaitingParticipant(request *model.AllowWaitingParticipantRequest) (*model.AllowWaitingParticipantResponse, error) {
	requestDef := GenReqDefForAllowWaitingParticipant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AllowWaitingParticipantResponse), nil
	}
}

// AllowWaitingParticipantInvoker 准入等候者
func (c *MeetingClient) AllowWaitingParticipantInvoker(request *model.AllowWaitingParticipantRequest) *AllowWaitingParticipantInvoker {
	requestDef := GenReqDefForAllowWaitingParticipant()
	return &AllowWaitingParticipantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateVmr 分配云会议室
//
// 企业管理员通过该接口将云会议室分配给用户、专业会议终端（TE10、TE20、HUAWEI Board、HUAWEI Bar 500及HUAWEI Box系列）、智慧屏TV、电子白板（SmartRooms）、IdeaHub。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) AssociateVmr(request *model.AssociateVmrRequest) (*model.AssociateVmrResponse, error) {
	requestDef := GenReqDefForAssociateVmr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateVmrResponse), nil
	}
}

// AssociateVmrInvoker 分配云会议室
func (c *MeetingClient) AssociateVmrInvoker(request *model.AssociateVmrRequest) *AssociateVmrInvoker {
	requestDef := GenReqDefForAssociateVmr()
	return &AssociateVmrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteCorpAdmins 批量删除企业管理员
//
// 通过该接口批量删除企业管理员。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) BatchDeleteCorpAdmins(request *model.BatchDeleteCorpAdminsRequest) (*model.BatchDeleteCorpAdminsResponse, error) {
	requestDef := GenReqDefForBatchDeleteCorpAdmins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteCorpAdminsResponse), nil
	}
}

// BatchDeleteCorpAdminsInvoker 批量删除企业管理员
func (c *MeetingClient) BatchDeleteCorpAdminsInvoker(request *model.BatchDeleteCorpAdminsRequest) *BatchDeleteCorpAdminsInvoker {
	requestDef := GenReqDefForBatchDeleteCorpAdmins()
	return &BatchDeleteCorpAdminsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteDevices 批量删除终端
//
// 企业管理员通过该接口批量删除专业会议终端，返回删除失败的列表。
// &gt; 如果需要删除Ideahub、SmartRooms、智慧屏TV请使用[[批量删除用户](https://support.huaweicloud.com/api-meeting/meeting_21_0070.html)](tag:hws)[[批量删除用户](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0070.html)](tag:hk)接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) BatchDeleteDevices(request *model.BatchDeleteDevicesRequest) (*model.BatchDeleteDevicesResponse, error) {
	requestDef := GenReqDefForBatchDeleteDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteDevicesResponse), nil
	}
}

// BatchDeleteDevicesInvoker 批量删除终端
func (c *MeetingClient) BatchDeleteDevicesInvoker(request *model.BatchDeleteDevicesRequest) *BatchDeleteDevicesInvoker {
	requestDef := GenReqDefForBatchDeleteDevices()
	return &BatchDeleteDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteMaterials 删除信息窗素材
//
// 删除信息窗素材。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) BatchDeleteMaterials(request *model.BatchDeleteMaterialsRequest) (*model.BatchDeleteMaterialsResponse, error) {
	requestDef := GenReqDefForBatchDeleteMaterials()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteMaterialsResponse), nil
	}
}

// BatchDeleteMaterialsInvoker 删除信息窗素材
func (c *MeetingClient) BatchDeleteMaterialsInvoker(request *model.BatchDeleteMaterialsRequest) *BatchDeleteMaterialsInvoker {
	requestDef := GenReqDefForBatchDeleteMaterials()
	return &BatchDeleteMaterialsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeletePrograms 删除信息窗节目
//
// 删除信息窗节目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) BatchDeletePrograms(request *model.BatchDeleteProgramsRequest) (*model.BatchDeleteProgramsResponse, error) {
	requestDef := GenReqDefForBatchDeletePrograms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteProgramsResponse), nil
	}
}

// BatchDeleteProgramsInvoker 删除信息窗节目
func (c *MeetingClient) BatchDeleteProgramsInvoker(request *model.BatchDeleteProgramsRequest) *BatchDeleteProgramsInvoker {
	requestDef := GenReqDefForBatchDeletePrograms()
	return &BatchDeleteProgramsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeletePublications 删除信息窗发布
//
// 删除信息窗发布。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) BatchDeletePublications(request *model.BatchDeletePublicationsRequest) (*model.BatchDeletePublicationsResponse, error) {
	requestDef := GenReqDefForBatchDeletePublications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeletePublicationsResponse), nil
	}
}

// BatchDeletePublicationsInvoker 删除信息窗发布
func (c *MeetingClient) BatchDeletePublicationsInvoker(request *model.BatchDeletePublicationsRequest) *BatchDeletePublicationsInvoker {
	requestDef := GenReqDefForBatchDeletePublications()
	return &BatchDeletePublicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteUsers 批量删除用户
//
// 企业管理员通过该接口批量删除企业用户。删除多个用户时，全部删除成功或者全部删除失败。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) BatchDeleteUsers(request *model.BatchDeleteUsersRequest) (*model.BatchDeleteUsersResponse, error) {
	requestDef := GenReqDefForBatchDeleteUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteUsersResponse), nil
	}
}

// BatchDeleteUsersInvoker 批量删除用户
func (c *MeetingClient) BatchDeleteUsersInvoker(request *model.BatchDeleteUsersRequest) *BatchDeleteUsersInvoker {
	requestDef := GenReqDefForBatchDeleteUsers()
	return &BatchDeleteUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchHand 批量举手
//
// 该接口用于批量设置来宾的举手/放下举手状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) BatchHand(request *model.BatchHandRequest) (*model.BatchHandResponse, error) {
	requestDef := GenReqDefForBatchHand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchHandResponse), nil
	}
}

// BatchHandInvoker 批量举手
func (c *MeetingClient) BatchHandInvoker(request *model.BatchHandRequest) *BatchHandInvoker {
	requestDef := GenReqDefForBatchHand()
	return &BatchHandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSearchAppId 分页查询企业应用
//
// 企业默认管理员分页查询企业应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) BatchSearchAppId(request *model.BatchSearchAppIdRequest) (*model.BatchSearchAppIdResponse, error) {
	requestDef := GenReqDefForBatchSearchAppId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSearchAppIdResponse), nil
	}
}

// BatchSearchAppIdInvoker 分页查询企业应用
func (c *MeetingClient) BatchSearchAppIdInvoker(request *model.BatchSearchAppIdRequest) *BatchSearchAppIdInvoker {
	requestDef := GenReqDefForBatchSearchAppId()
	return &BatchSearchAppIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowUserDetails 批量查询用户详情
//
// 批量查询用户详情，支持指定第三方账号查询详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) BatchShowUserDetails(request *model.BatchShowUserDetailsRequest) (*model.BatchShowUserDetailsResponse, error) {
	requestDef := GenReqDefForBatchShowUserDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowUserDetailsResponse), nil
	}
}

// BatchShowUserDetailsInvoker 批量查询用户详情
func (c *MeetingClient) BatchShowUserDetailsInvoker(request *model.BatchShowUserDetailsRequest) *BatchShowUserDetailsInvoker {
	requestDef := GenReqDefForBatchShowUserDetails()
	return &BatchShowUserDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateDevicesStatus 批量修改终端状态
//
// 企业管理员通过该接口批量修改专业会议终端状态。当硬终端资源到期后，若企业内对应资源的硬终端超过数量后会被系统随机自动停用，此时可通过该接口修改专业会议终端的状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) BatchUpdateDevicesStatus(request *model.BatchUpdateDevicesStatusRequest) (*model.BatchUpdateDevicesStatusResponse, error) {
	requestDef := GenReqDefForBatchUpdateDevicesStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateDevicesStatusResponse), nil
	}
}

// BatchUpdateDevicesStatusInvoker 批量修改终端状态
func (c *MeetingClient) BatchUpdateDevicesStatusInvoker(request *model.BatchUpdateDevicesStatusRequest) *BatchUpdateDevicesStatusInvoker {
	requestDef := GenReqDefForBatchUpdateDevicesStatus()
	return &BatchUpdateDevicesStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateUserStatus 批量修改用户状态
//
// 企业管理员通过该接口批量修改用户状态，当用户帐号数资源或者电子白板（SmartRooms）资源到期后，若企业内对应资源的用户帐号超过数量后会被系统随机自动停用，此时可通过该接口修改用户的状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) BatchUpdateUserStatus(request *model.BatchUpdateUserStatusRequest) (*model.BatchUpdateUserStatusResponse, error) {
	requestDef := GenReqDefForBatchUpdateUserStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateUserStatusResponse), nil
	}
}

// BatchUpdateUserStatusInvoker 批量修改用户状态
func (c *MeetingClient) BatchUpdateUserStatusInvoker(request *model.BatchUpdateUserStatusRequest) *BatchUpdateUserStatusInvoker {
	requestDef := GenReqDefForBatchUpdateUserStatus()
	return &BatchUpdateUserStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BroadcastParticipant 广播会场
//
// 该接口用于广播指定的与会者。同一时间，只允许一个与会者被广播。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) BroadcastParticipant(request *model.BroadcastParticipantRequest) (*model.BroadcastParticipantResponse, error) {
	requestDef := GenReqDefForBroadcastParticipant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BroadcastParticipantResponse), nil
	}
}

// BroadcastParticipantInvoker 广播会场
func (c *MeetingClient) BroadcastParticipantInvoker(request *model.BroadcastParticipantRequest) *BroadcastParticipantInvoker {
	requestDef := GenReqDefForBroadcastParticipant()
	return &BroadcastParticipantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelBroadcast 取消广播
//
// 该接口用于取消广播，包括：取消广播多画面，取消广播会场，取消点名会场。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CancelBroadcast(request *model.CancelBroadcastRequest) (*model.CancelBroadcastResponse, error) {
	requestDef := GenReqDefForCancelBroadcast()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelBroadcastResponse), nil
	}
}

// CancelBroadcastInvoker 取消广播
func (c *MeetingClient) CancelBroadcastInvoker(request *model.CancelBroadcastRequest) *CancelBroadcastInvoker {
	requestDef := GenReqDefForCancelBroadcast()
	return &CancelBroadcastInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelMeeting 取消预约会议
//
// 该接口用于取消预约的会议。企业管理员可以取消本企业下用户创建的会议，普通用户只能取消自己创建的会议。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CancelMeeting(request *model.CancelMeetingRequest) (*model.CancelMeetingResponse, error) {
	requestDef := GenReqDefForCancelMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelMeetingResponse), nil
	}
}

// CancelMeetingInvoker 取消预约会议
func (c *MeetingClient) CancelMeetingInvoker(request *model.CancelMeetingRequest) *CancelMeetingInvoker {
	requestDef := GenReqDefForCancelMeeting()
	return &CancelMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelRecurringMeeting 取消周期性会议
//
// 该接口用于取消周期性会议。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CancelRecurringMeeting(request *model.CancelRecurringMeetingRequest) (*model.CancelRecurringMeetingResponse, error) {
	requestDef := GenReqDefForCancelRecurringMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelRecurringMeetingResponse), nil
	}
}

// CancelRecurringMeetingInvoker 取消周期性会议
func (c *MeetingClient) CancelRecurringMeetingInvoker(request *model.CancelRecurringMeetingRequest) *CancelRecurringMeetingInvoker {
	requestDef := GenReqDefForCancelRecurringMeeting()
	return &CancelRecurringMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelRecurringSubMeeting 取消周期性会议的子会议
//
// 该接口用于取消周期性会议的子会议。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CancelRecurringSubMeeting(request *model.CancelRecurringSubMeetingRequest) (*model.CancelRecurringSubMeetingResponse, error) {
	requestDef := GenReqDefForCancelRecurringSubMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelRecurringSubMeetingResponse), nil
	}
}

// CancelRecurringSubMeetingInvoker 取消周期性会议的子会议
func (c *MeetingClient) CancelRecurringSubMeetingInvoker(request *model.CancelRecurringSubMeetingRequest) *CancelRecurringSubMeetingInvoker {
	requestDef := GenReqDefForCancelRecurringSubMeeting()
	return &CancelRecurringSubMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckSlideVerifyCode 校验滑块验证码
//
// 该接口提供校验滑块验证码。服务器收到请求，返回校验结果。用户在前台界面通过滑块操作匹配图形，使得抠图和原图吻合。然后服务器进行校验滑块验证码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CheckSlideVerifyCode(request *model.CheckSlideVerifyCodeRequest) (*model.CheckSlideVerifyCodeResponse, error) {
	requestDef := GenReqDefForCheckSlideVerifyCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckSlideVerifyCodeResponse), nil
	}
}

// CheckSlideVerifyCodeInvoker 校验滑块验证码
func (c *MeetingClient) CheckSlideVerifyCodeInvoker(request *model.CheckSlideVerifyCodeRequest) *CheckSlideVerifyCodeInvoker {
	requestDef := GenReqDefForCheckSlideVerifyCode()
	return &CheckSlideVerifyCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckToken 校验Token
//
// 该接口提供校验token合法性功能。服务器收到请求后，验证token合法性并返回结果。如果参数needGenNewToken为true时，生成新的token并返回。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CheckToken(request *model.CheckTokenRequest) (*model.CheckTokenResponse, error) {
	requestDef := GenReqDefForCheckToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckTokenResponse), nil
	}
}

// CheckTokenInvoker 校验Token
func (c *MeetingClient) CheckTokenInvoker(request *model.CheckTokenRequest) *CheckTokenInvoker {
	requestDef := GenReqDefForCheckToken()
	return &CheckTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckVeriCodeForUpdateUserInfo 校验手机和邮箱对应的验证码
//
// 企业用户通过该接口校验手机和邮箱对应的验证码，一分钟内记录尝试次数不得超过5次。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CheckVeriCodeForUpdateUserInfo(request *model.CheckVeriCodeForUpdateUserInfoRequest) (*model.CheckVeriCodeForUpdateUserInfoResponse, error) {
	requestDef := GenReqDefForCheckVeriCodeForUpdateUserInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckVeriCodeForUpdateUserInfoResponse), nil
	}
}

// CheckVeriCodeForUpdateUserInfoInvoker 校验手机和邮箱对应的验证码
func (c *MeetingClient) CheckVeriCodeForUpdateUserInfoInvoker(request *model.CheckVeriCodeForUpdateUserInfoRequest) *CheckVeriCodeForUpdateUserInfoInvoker {
	requestDef := GenReqDefForCheckVeriCodeForUpdateUserInfo()
	return &CheckVeriCodeForUpdateUserInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckVerifyCode 校验验证码
//
// 该接口提供校验验证码，服务器收到请求，返回结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CheckVerifyCode(request *model.CheckVerifyCodeRequest) (*model.CheckVerifyCodeResponse, error) {
	requestDef := GenReqDefForCheckVerifyCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckVerifyCodeResponse), nil
	}
}

// CheckVerifyCodeInvoker 校验验证码
func (c *MeetingClient) CheckVerifyCodeInvoker(request *model.CheckVerifyCodeRequest) *CheckVerifyCodeInvoker {
	requestDef := GenReqDefForCheckVerifyCode()
	return &CheckVerifyCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAnonymousAuthRandom 匿名用户会议鉴权
//
// 该接口用于匿名用户入会鉴权。请求根据会议ID和密码鉴权，返回鉴权随机数（可以根据该随机数获取匿名用户信息、会议信息等）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CreateAnonymousAuthRandom(request *model.CreateAnonymousAuthRandomRequest) (*model.CreateAnonymousAuthRandomResponse, error) {
	requestDef := GenReqDefForCreateAnonymousAuthRandom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAnonymousAuthRandomResponse), nil
	}
}

// CreateAnonymousAuthRandomInvoker 匿名用户会议鉴权
func (c *MeetingClient) CreateAnonymousAuthRandomInvoker(request *model.CreateAnonymousAuthRandomRequest) *CreateAnonymousAuthRandomInvoker {
	requestDef := GenReqDefForCreateAnonymousAuthRandom()
	return &CreateAnonymousAuthRandomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuthRandom 获取会议鉴权随机数
//
// 根据会议ID + 密码鉴权返回鉴权随机数，如果是小程序调用时，需要企业支持小程序功能
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CreateAuthRandom(request *model.CreateAuthRandomRequest) (*model.CreateAuthRandomResponse, error) {
	requestDef := GenReqDefForCreateAuthRandom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuthRandomResponse), nil
	}
}

// CreateAuthRandomInvoker 获取会议鉴权随机数
func (c *MeetingClient) CreateAuthRandomInvoker(request *model.CreateAuthRandomRequest) *CreateAuthRandomInvoker {
	requestDef := GenReqDefForCreateAuthRandom()
	return &CreateAuthRandomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConfToken 获取会控Token
//
// 该接口用于获取正在召开会议的会控Token（未开始的会议调用该接口返回失败）。Token有效期是半个小时。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CreateConfToken(request *model.CreateConfTokenRequest) (*model.CreateConfTokenResponse, error) {
	requestDef := GenReqDefForCreateConfToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConfTokenResponse), nil
	}
}

// CreateConfTokenInvoker 获取会控Token
func (c *MeetingClient) CreateConfTokenInvoker(request *model.CreateConfTokenRequest) *CreateConfTokenInvoker {
	requestDef := GenReqDefForCreateConfToken()
	return &CreateConfTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMeeting 创建会议
//
// 该接口用于创建立即会议和预约会议。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CreateMeeting(request *model.CreateMeetingRequest) (*model.CreateMeetingResponse, error) {
	requestDef := GenReqDefForCreateMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMeetingResponse), nil
	}
}

// CreateMeetingInvoker 创建会议
func (c *MeetingClient) CreateMeetingInvoker(request *model.CreateMeetingRequest) *CreateMeetingInvoker {
	requestDef := GenReqDefForCreateMeeting()
	return &CreateMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePortalRefNonce 获取页面免登陆跳转的nonce信息
//
// 通过Access Token生成页面免登陆跳转到华为云会议的Portal的nonce信息。获取到nonce信息后，通过链接https://meeting.huaweicloud.com/?lang&#x3D;zh-CN&amp;nonce&#x3D;xxxxxxxxxxxxx#/login进行免登陆跳转。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CreatePortalRefNonce(request *model.CreatePortalRefNonceRequest) (*model.CreatePortalRefNonceResponse, error) {
	requestDef := GenReqDefForCreatePortalRefNonce()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePortalRefNonceResponse), nil
	}
}

// CreatePortalRefNonceInvoker 获取页面免登陆跳转的nonce信息
func (c *MeetingClient) CreatePortalRefNonceInvoker(request *model.CreatePortalRefNonceRequest) *CreatePortalRefNonceInvoker {
	requestDef := GenReqDefForCreatePortalRefNonce()
	return &CreatePortalRefNonceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRecurringMeeting 创建周期性会议
//
// 该接口用于预约周期性会议。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CreateRecurringMeeting(request *model.CreateRecurringMeetingRequest) (*model.CreateRecurringMeetingResponse, error) {
	requestDef := GenReqDefForCreateRecurringMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecurringMeetingResponse), nil
	}
}

// CreateRecurringMeetingInvoker 创建周期性会议
func (c *MeetingClient) CreateRecurringMeetingInvoker(request *model.CreateRecurringMeetingRequest) *CreateRecurringMeetingInvoker {
	requestDef := GenReqDefForCreateRecurringMeeting()
	return &CreateRecurringMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVisionActiveCode 企业管理员生成激活码
//
// 企业管理员生成智慧屏、电子白板（SmartRooms）、Ideahub的激活码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CreateVisionActiveCode(request *model.CreateVisionActiveCodeRequest) (*model.CreateVisionActiveCodeResponse, error) {
	requestDef := GenReqDefForCreateVisionActiveCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVisionActiveCodeResponse), nil
	}
}

// CreateVisionActiveCodeInvoker 企业管理员生成激活码
func (c *MeetingClient) CreateVisionActiveCodeInvoker(request *model.CreateVisionActiveCodeRequest) *CreateVisionActiveCodeInvoker {
	requestDef := GenReqDefForCreateVisionActiveCode()
	return &CreateVisionActiveCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWebSocketToken 获取websocket建链Token
//
// 该接口用于获取会控WebSocket建链的临时Token。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CreateWebSocketToken(request *model.CreateWebSocketTokenRequest) (*model.CreateWebSocketTokenResponse, error) {
	requestDef := GenReqDefForCreateWebSocketToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWebSocketTokenResponse), nil
	}
}

// CreateWebSocketTokenInvoker 获取websocket建链Token
func (c *MeetingClient) CreateWebSocketTokenInvoker(request *model.CreateWebSocketTokenRequest) *CreateWebSocketTokenInvoker {
	requestDef := GenReqDefForCreateWebSocketToken()
	return &CreateWebSocketTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWebinar 预约网络研讨会
//
// 该接口用于创建网络研讨会。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) CreateWebinar(request *model.CreateWebinarRequest) (*model.CreateWebinarResponse, error) {
	requestDef := GenReqDefForCreateWebinar()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWebinarResponse), nil
	}
}

// CreateWebinarInvoker 预约网络研讨会
func (c *MeetingClient) CreateWebinarInvoker(request *model.CreateWebinarRequest) *CreateWebinarInvoker {
	requestDef := GenReqDefForCreateWebinar()
	return &CreateWebinarInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppId 删除企业应用
//
// 企业管理员删除企业应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) DeleteAppId(request *model.DeleteAppIdRequest) (*model.DeleteAppIdResponse, error) {
	requestDef := GenReqDefForDeleteAppId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppIdResponse), nil
	}
}

// DeleteAppIdInvoker 删除企业应用
func (c *MeetingClient) DeleteAppIdInvoker(request *model.DeleteAppIdRequest) *DeleteAppIdInvoker {
	requestDef := GenReqDefForDeleteAppId()
	return &DeleteAppIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAttendees 删除与会者
//
// 该接口用于删除与会者。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) DeleteAttendees(request *model.DeleteAttendeesRequest) (*model.DeleteAttendeesResponse, error) {
	requestDef := GenReqDefForDeleteAttendees()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAttendeesResponse), nil
	}
}

// DeleteAttendeesInvoker 删除与会者
func (c *MeetingClient) DeleteAttendeesInvoker(request *model.DeleteAttendeesRequest) *DeleteAttendeesInvoker {
	requestDef := GenReqDefForDeleteAttendees()
	return &DeleteAttendeesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCorp SP管理员删除企业
//
// 删除企业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) DeleteCorp(request *model.DeleteCorpRequest) (*model.DeleteCorpResponse, error) {
	requestDef := GenReqDefForDeleteCorp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCorpResponse), nil
	}
}

// DeleteCorpInvoker SP管理员删除企业
func (c *MeetingClient) DeleteCorpInvoker(request *model.DeleteCorpRequest) *DeleteCorpInvoker {
	requestDef := GenReqDefForDeleteCorp()
	return &DeleteCorpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCorpVmr 删除云会议室
//
// 企业管理员通过该接口删除企业的云会议室。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) DeleteCorpVmr(request *model.DeleteCorpVmrRequest) (*model.DeleteCorpVmrResponse, error) {
	requestDef := GenReqDefForDeleteCorpVmr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCorpVmrResponse), nil
	}
}

// DeleteCorpVmrInvoker 删除云会议室
func (c *MeetingClient) DeleteCorpVmrInvoker(request *model.DeleteCorpVmrRequest) *DeleteCorpVmrInvoker {
	requestDef := GenReqDefForDeleteCorpVmr()
	return &DeleteCorpVmrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDepartment 删除部门
//
// 企业管理员通过该接口删除部门。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) DeleteDepartment(request *model.DeleteDepartmentRequest) (*model.DeleteDepartmentResponse, error) {
	requestDef := GenReqDefForDeleteDepartment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDepartmentResponse), nil
	}
}

// DeleteDepartmentInvoker 删除部门
func (c *MeetingClient) DeleteDepartmentInvoker(request *model.DeleteDepartmentRequest) *DeleteDepartmentInvoker {
	requestDef := GenReqDefForDeleteDepartment()
	return &DeleteDepartmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLayout 删除多画面布局
//
// 该接口用于删除当前会议已保存的多画面布局。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) DeleteLayout(request *model.DeleteLayoutRequest) (*model.DeleteLayoutResponse, error) {
	requestDef := GenReqDefForDeleteLayout()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLayoutResponse), nil
	}
}

// DeleteLayoutInvoker 删除多画面布局
func (c *MeetingClient) DeleteLayoutInvoker(request *model.DeleteLayoutRequest) *DeleteLayoutInvoker {
	requestDef := GenReqDefForDeleteLayout()
	return &DeleteLayoutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRecordings 批量删除录制
//
// 该接口用于批量删除会议的录制。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) DeleteRecordings(request *model.DeleteRecordingsRequest) (*model.DeleteRecordingsResponse, error) {
	requestDef := GenReqDefForDeleteRecordings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRecordingsResponse), nil
	}
}

// DeleteRecordingsInvoker 批量删除录制
func (c *MeetingClient) DeleteRecordingsInvoker(request *model.DeleteRecordingsRequest) *DeleteRecordingsInvoker {
	requestDef := GenReqDefForDeleteRecordings()
	return &DeleteRecordingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResource SP管理员根据删除企业资源
//
// 企业删除资源项，删除资源项后，企业资源总数会自动减少。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) DeleteResource(request *model.DeleteResourceRequest) (*model.DeleteResourceResponse, error) {
	requestDef := GenReqDefForDeleteResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceResponse), nil
	}
}

// DeleteResourceInvoker SP管理员根据删除企业资源
func (c *MeetingClient) DeleteResourceInvoker(request *model.DeleteResourceRequest) *DeleteResourceInvoker {
	requestDef := GenReqDefForDeleteResource()
	return &DeleteResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteToken 注销登录
//
// 该接口提供注销功能。服务器收到请求后，删除该Token。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) DeleteToken(request *model.DeleteTokenRequest) (*model.DeleteTokenResponse, error) {
	requestDef := GenReqDefForDeleteToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTokenResponse), nil
	}
}

// DeleteTokenInvoker 注销登录
func (c *MeetingClient) DeleteTokenInvoker(request *model.DeleteTokenRequest) *DeleteTokenInvoker {
	requestDef := GenReqDefForDeleteToken()
	return &DeleteTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVisionActiveCode 企业管理员删除激活码
//
// 企业管理员批量删除激活码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) DeleteVisionActiveCode(request *model.DeleteVisionActiveCodeRequest) (*model.DeleteVisionActiveCodeResponse, error) {
	requestDef := GenReqDefForDeleteVisionActiveCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVisionActiveCodeResponse), nil
	}
}

// DeleteVisionActiveCodeInvoker 企业管理员删除激活码
func (c *MeetingClient) DeleteVisionActiveCodeInvoker(request *model.DeleteVisionActiveCodeRequest) *DeleteVisionActiveCodeInvoker {
	requestDef := GenReqDefForDeleteVisionActiveCode()
	return &DeleteVisionActiveCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWebHookConfig 删除事件推送
//
// 该接口用于管理员删除已配置的事件推送设置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) DeleteWebHookConfig(request *model.DeleteWebHookConfigRequest) (*model.DeleteWebHookConfigResponse, error) {
	requestDef := GenReqDefForDeleteWebHookConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWebHookConfigResponse), nil
	}
}

// DeleteWebHookConfigInvoker 删除事件推送
func (c *MeetingClient) DeleteWebHookConfigInvoker(request *model.DeleteWebHookConfigRequest) *DeleteWebHookConfigInvoker {
	requestDef := GenReqDefForDeleteWebHookConfig()
	return &DeleteWebHookConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWebinar 取消网络研讨会
//
// 该接口用于取消已预约的网络研讨会。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) DeleteWebinar(request *model.DeleteWebinarRequest) (*model.DeleteWebinarResponse, error) {
	requestDef := GenReqDefForDeleteWebinar()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWebinarResponse), nil
	}
}

// DeleteWebinarInvoker 取消网络研讨会
func (c *MeetingClient) DeleteWebinarInvoker(request *model.DeleteWebinarRequest) *DeleteWebinarInvoker {
	requestDef := GenReqDefForDeleteWebinar()
	return &DeleteWebinarInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateVmr 回收云会议室
//
// 企业管理员通过该接口回收云会议室。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) DisassociateVmr(request *model.DisassociateVmrRequest) (*model.DisassociateVmrResponse, error) {
	requestDef := GenReqDefForDisassociateVmr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateVmrResponse), nil
	}
}

// DisassociateVmrInvoker 回收云会议室
func (c *MeetingClient) DisassociateVmrInvoker(request *model.DisassociateVmrRequest) *DisassociateVmrInvoker {
	requestDef := GenReqDefForDisassociateVmr()
	return &DisassociateVmrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Hand 举手
//
// 该接口用于设置指定与会者的举手/放下举手状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) Hand(request *model.HandRequest) (*model.HandResponse, error) {
	requestDef := GenReqDefForHand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.HandResponse), nil
	}
}

// HandInvoker 举手
func (c *MeetingClient) HandInvoker(request *model.HandRequest) *HandInvoker {
	requestDef := GenReqDefForHand()
	return &HandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// HangUp 挂断与会者
//
// 该接口用于挂断正在通话中的与会者。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) HangUp(request *model.HangUpRequest) (*model.HangUpResponse, error) {
	requestDef := GenReqDefForHangUp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.HangUpResponse), nil
	}
}

// HangUpInvoker 挂断与会者
func (c *MeetingClient) HangUpInvoker(request *model.HangUpRequest) *HangUpInvoker {
	requestDef := GenReqDefForHangUp()
	return &HangUpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InviteOperateVideo 主持人邀请与会者开启/关闭摄像头
//
// 该接口用于邀请指定与会者开启、关闭摄像头。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) InviteOperateVideo(request *model.InviteOperateVideoRequest) (*model.InviteOperateVideoResponse, error) {
	requestDef := GenReqDefForInviteOperateVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InviteOperateVideoResponse), nil
	}
}

// InviteOperateVideoInvoker 主持人邀请与会者开启/关闭摄像头
func (c *MeetingClient) InviteOperateVideoInvoker(request *model.InviteOperateVideoRequest) *InviteOperateVideoInvoker {
	requestDef := GenReqDefForInviteOperateVideo()
	return &InviteOperateVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InviteParticipant 邀请与会者
//
// 该接口用于邀请与会者加入会议。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) InviteParticipant(request *model.InviteParticipantRequest) (*model.InviteParticipantResponse, error) {
	requestDef := GenReqDefForInviteParticipant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InviteParticipantResponse), nil
	}
}

// InviteParticipantInvoker 邀请与会者
func (c *MeetingClient) InviteParticipantInvoker(request *model.InviteParticipantRequest) *InviteParticipantInvoker {
	requestDef := GenReqDefForInviteParticipant()
	return &InviteParticipantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InviteShare 邀请共享
//
// 该接口用于邀请/取消邀请指定与会人共享桌面。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) InviteShare(request *model.InviteShareRequest) (*model.InviteShareResponse, error) {
	requestDef := GenReqDefForInviteShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InviteShareResponse), nil
	}
}

// InviteShareInvoker 邀请共享
func (c *MeetingClient) InviteShareInvoker(request *model.InviteShareRequest) *InviteShareInvoker {
	requestDef := GenReqDefForInviteShare()
	return &InviteShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InviteUser 邀请用户
//
// 通过手机号码或者邮箱地址邀请用户加入企业。
// * 若被邀请用户在华为云会议系统中不存在，则：
//   - 华为云会议免费版和华为云会议标准版发送短信/邮件邀请用户完成注册后加入企业。用户注册成功后，加入该企业。
//   - 华为云会议旗舰版在企业内直接添加该用户。用户会收到华为云会议的初始密码，用户第一次以手机号或者邮箱登录时，需要修改密码。
//
// * 若被邀请用户在华为云会议系统中存在，则该用户会收到短信或者邮件确认。确认完成后改用户加入企业内。该用户的密码保持原来的密码不变。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) InviteUser(request *model.InviteUserRequest) (*model.InviteUserResponse, error) {
	requestDef := GenReqDefForInviteUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InviteUserResponse), nil
	}
}

// InviteUserInvoker 邀请用户
func (c *MeetingClient) InviteUserInvoker(request *model.InviteUserRequest) *InviteUserInvoker {
	requestDef := GenReqDefForInviteUser()
	return &InviteUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InviteWithPwd 通过会议ID和密码邀请与会者
//
// 该接口用于通过会议ID和密码邀请与会者。一般用于App已知会议ID和来宾密码，通过扫码等方式获取其他终端的SIP号码后，使用该接口将其他终端邀请加入会议中。
// &gt; 需要管理员在企业的“会议设置”&gt;“来宾扫码邀请任意硬终端入会”设置成打开，才允许通过来宾密码邀请其他终端入会。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) InviteWithPwd(request *model.InviteWithPwdRequest) (*model.InviteWithPwdResponse, error) {
	requestDef := GenReqDefForInviteWithPwd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InviteWithPwdResponse), nil
	}
}

// InviteWithPwdInvoker 通过会议ID和密码邀请与会者
func (c *MeetingClient) InviteWithPwdInvoker(request *model.InviteWithPwdRequest) *InviteWithPwdInvoker {
	requestDef := GenReqDefForInviteWithPwd()
	return &InviteWithPwdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHistoryWebinars 查询历史召开的网络研讨会列表
//
// 该接口用于查询历史网络研讨会。管理员可查询企业内历史网络研讨会，非管理员可查询个人历史网络研讨会。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ListHistoryWebinars(request *model.ListHistoryWebinarsRequest) (*model.ListHistoryWebinarsResponse, error) {
	requestDef := GenReqDefForListHistoryWebinars()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryWebinarsResponse), nil
	}
}

// ListHistoryWebinarsInvoker 查询历史召开的网络研讨会列表
func (c *MeetingClient) ListHistoryWebinarsInvoker(request *model.ListHistoryWebinarsRequest) *ListHistoryWebinarsInvoker {
	requestDef := GenReqDefForListHistoryWebinars()
	return &ListHistoryWebinarsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNetworkQuality 查询会场网络质量
//
// 查询会场网络质量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ListNetworkQuality(request *model.ListNetworkQualityRequest) (*model.ListNetworkQualityResponse, error) {
	requestDef := GenReqDefForListNetworkQuality()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNetworkQualityResponse), nil
	}
}

// ListNetworkQualityInvoker 查询会场网络质量
func (c *MeetingClient) ListNetworkQualityInvoker(request *model.ListNetworkQualityRequest) *ListNetworkQualityInvoker {
	requestDef := GenReqDefForListNetworkQuality()
	return &ListNetworkQualityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOngoingWebinars 查询正在召开的网络研讨会列表
//
// 该接口用于查询正在召开的网络研讨会。管理员可查询企业内正在召开网络研讨会，非管理员可查询自己预订的正在召开的网络研讨会。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ListOngoingWebinars(request *model.ListOngoingWebinarsRequest) (*model.ListOngoingWebinarsResponse, error) {
	requestDef := GenReqDefForListOngoingWebinars()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOngoingWebinarsResponse), nil
	}
}

// ListOngoingWebinarsInvoker 查询正在召开的网络研讨会列表
func (c *MeetingClient) ListOngoingWebinarsInvoker(request *model.ListOngoingWebinarsRequest) *ListOngoingWebinarsInvoker {
	requestDef := GenReqDefForListOngoingWebinars()
	return &ListOngoingWebinarsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUpComingWebinars 查询即将召开的网络研讨会列表
//
// 该接口用于查询即将召开的网络研讨会。管理员可查询企业内即将召开网络研讨会，非管理员可查询自己预订的即将召开的网络研讨会。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ListUpComingWebinars(request *model.ListUpComingWebinarsRequest) (*model.ListUpComingWebinarsResponse, error) {
	requestDef := GenReqDefForListUpComingWebinars()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUpComingWebinarsResponse), nil
	}
}

// ListUpComingWebinarsInvoker 查询即将召开的网络研讨会列表
func (c *MeetingClient) ListUpComingWebinarsInvoker(request *model.ListUpComingWebinarsRequest) *ListUpComingWebinarsInvoker {
	requestDef := GenReqDefForListUpComingWebinars()
	return &ListUpComingWebinarsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Live 启停会议直播
//
// 该接口用于启动或停止会议直播。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) Live(request *model.LiveRequest) (*model.LiveResponse, error) {
	requestDef := GenReqDefForLive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LiveResponse), nil
	}
}

// LiveInvoker 启停会议直播
func (c *MeetingClient) LiveInvoker(request *model.LiveRequest) *LiveInvoker {
	requestDef := GenReqDefForLive()
	return &LiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LockMeeting 锁定会议
//
// 该接口用于锁定或解锁会议。锁定会议后，不允许新的来宾主动加入会议。会议锁定后使用主持人密码/主持人链接加入会议或者主持人邀请来宾不受影响。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) LockMeeting(request *model.LockMeetingRequest) (*model.LockMeetingResponse, error) {
	requestDef := GenReqDefForLockMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LockMeetingResponse), nil
	}
}

// LockMeetingInvoker 锁定会议
func (c *MeetingClient) LockMeetingInvoker(request *model.LockMeetingRequest) *LockMeetingInvoker {
	requestDef := GenReqDefForLockMeeting()
	return &LockMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LockView 锁定会场视频源
//
// 该接口用于锁定或者解锁某在线会场的视频源。只适用于专业会议终端（如TE系列等）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) LockView(request *model.LockViewRequest) (*model.LockViewResponse, error) {
	requestDef := GenReqDefForLockView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LockViewResponse), nil
	}
}

// LockViewInvoker 锁定会场视频源
func (c *MeetingClient) LockViewInvoker(request *model.LockViewRequest) *LockViewInvoker {
	requestDef := GenReqDefForLockView()
	return &LockViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MoveToWaitingRoom 移入等候室
//
// 该接口用于将会中的指定与会者移入到等候室。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) MoveToWaitingRoom(request *model.MoveToWaitingRoomRequest) (*model.MoveToWaitingRoomResponse, error) {
	requestDef := GenReqDefForMoveToWaitingRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MoveToWaitingRoomResponse), nil
	}
}

// MoveToWaitingRoomInvoker 移入等候室
func (c *MeetingClient) MoveToWaitingRoomInvoker(request *model.MoveToWaitingRoomRequest) *MoveToWaitingRoomInvoker {
	requestDef := GenReqDefForMoveToWaitingRoom()
	return &MoveToWaitingRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MuteMeeting 全场静音
//
// 该接口用于设置整个会议所有与会者（主持人除外）的静音/取消静音状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) MuteMeeting(request *model.MuteMeetingRequest) (*model.MuteMeetingResponse, error) {
	requestDef := GenReqDefForMuteMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MuteMeetingResponse), nil
	}
}

// MuteMeetingInvoker 全场静音
func (c *MeetingClient) MuteMeetingInvoker(request *model.MuteMeetingRequest) *MuteMeetingInvoker {
	requestDef := GenReqDefForMuteMeeting()
	return &MuteMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MuteParticipant 静音与会者
//
// 该接口用于设置指定与会者静音/取消静音状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) MuteParticipant(request *model.MuteParticipantRequest) (*model.MuteParticipantResponse, error) {
	requestDef := GenReqDefForMuteParticipant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MuteParticipantResponse), nil
	}
}

// MuteParticipantInvoker 静音与会者
func (c *MeetingClient) MuteParticipantInvoker(request *model.MuteParticipantRequest) *MuteParticipantInvoker {
	requestDef := GenReqDefForMuteParticipant()
	return &MuteParticipantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ProlongMeeting 延长会议
//
// 该接口用于延长会议时间。默认会议自动延长。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ProlongMeeting(request *model.ProlongMeetingRequest) (*model.ProlongMeetingResponse, error) {
	requestDef := GenReqDefForProlongMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ProlongMeetingResponse), nil
	}
}

// ProlongMeetingInvoker 延长会议
func (c *MeetingClient) ProlongMeetingInvoker(request *model.ProlongMeetingRequest) *ProlongMeetingInvoker {
	requestDef := GenReqDefForProlongMeeting()
	return &ProlongMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Record 启停会议录制
//
// 该接口用于启动或停止会议云录制。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) Record(request *model.RecordRequest) (*model.RecordResponse, error) {
	requestDef := GenReqDefForRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecordResponse), nil
	}
}

// RecordInvoker 启停会议录制
func (c *MeetingClient) RecordInvoker(request *model.RecordRequest) *RecordInvoker {
	requestDef := GenReqDefForRecord()
	return &RecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RenameParticipant 重命名与会者
//
// 重命名某个与会者。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) RenameParticipant(request *model.RenameParticipantRequest) (*model.RenameParticipantResponse, error) {
	requestDef := GenReqDefForRenameParticipant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RenameParticipantResponse), nil
	}
}

// RenameParticipantInvoker 重命名与会者
func (c *MeetingClient) RenameParticipantInvoker(request *model.RenameParticipantRequest) *RenameParticipantInvoker {
	requestDef := GenReqDefForRenameParticipant()
	return &RenameParticipantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetActivecode 企业管理员通过sn重置激活码
//
// 当硬终端激活码失效时，企业管理员可以通过该接口重置激活码，使用重新获取的激活码激活终端，每24小时可重新激活5次。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ResetActivecode(request *model.ResetActivecodeRequest) (*model.ResetActivecodeResponse, error) {
	requestDef := GenReqDefForResetActivecode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetActivecodeResponse), nil
	}
}

// ResetActivecodeInvoker 企业管理员通过sn重置激活码
func (c *MeetingClient) ResetActivecodeInvoker(request *model.ResetActivecodeRequest) *ResetActivecodeInvoker {
	requestDef := GenReqDefForResetActivecode()
	return &ResetActivecodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetAppKey 重置企业应用appkey
//
// 企业默认管理员重置企业应用appkey
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ResetAppKey(request *model.ResetAppKeyRequest) (*model.ResetAppKeyResponse, error) {
	requestDef := GenReqDefForResetAppKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetAppKeyResponse), nil
	}
}

// ResetAppKeyInvoker 重置企业应用appkey
func (c *MeetingClient) ResetAppKeyInvoker(request *model.ResetAppKeyRequest) *ResetAppKeyInvoker {
	requestDef := GenReqDefForResetAppKey()
	return &ResetAppKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPwd 用户重置密码
//
// 该接口提供给用户重置密码功能，服务器收到请求，重新设置用户密码并返回结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ResetPwd(request *model.ResetPwdRequest) (*model.ResetPwdResponse, error) {
	requestDef := GenReqDefForResetPwd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPwdResponse), nil
	}
}

// ResetPwdInvoker 用户重置密码
func (c *MeetingClient) ResetPwdInvoker(request *model.ResetPwdRequest) *ResetPwdInvoker {
	requestDef := GenReqDefForResetPwd()
	return &ResetPwdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPwdByAdmin 企业管理员重置企业成员密码
//
// 企业管理员通过该接口提供企业管理员重置企业成员密码的功能。当服务器收到重置密码的请求时，发送新的密码到企业成员的邮箱或者短信，并返回结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ResetPwdByAdmin(request *model.ResetPwdByAdminRequest) (*model.ResetPwdByAdminResponse, error) {
	requestDef := GenReqDefForResetPwdByAdmin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPwdByAdminResponse), nil
	}
}

// ResetPwdByAdminInvoker 企业管理员重置企业成员密码
func (c *MeetingClient) ResetPwdByAdminInvoker(request *model.ResetPwdByAdminRequest) *ResetPwdByAdminInvoker {
	requestDef := GenReqDefForResetPwdByAdmin()
	return &ResetPwdByAdminInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetVisionActiveCode 企业管理员重置帐号的激活码
//
// 企业管理员重置帐号的激活码，重置后，原设备直接解绑，必须重新激活使用,若手机邮箱不填，则不会发送新的激活码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ResetVisionActiveCode(request *model.ResetVisionActiveCodeRequest) (*model.ResetVisionActiveCodeResponse, error) {
	requestDef := GenReqDefForResetVisionActiveCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetVisionActiveCodeResponse), nil
	}
}

// ResetVisionActiveCodeInvoker 企业管理员重置帐号的激活码
func (c *MeetingClient) ResetVisionActiveCodeInvoker(request *model.ResetVisionActiveCodeRequest) *ResetVisionActiveCodeInvoker {
	requestDef := GenReqDefForResetVisionActiveCode()
	return &ResetVisionActiveCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResumeSimultaneousInterpretation 开启/关闭同声传译
//
// 该接口用于会议主席可以通过该接口开启/关闭同声传译。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ResumeSimultaneousInterpretation(request *model.ResumeSimultaneousInterpretationRequest) (*model.ResumeSimultaneousInterpretationResponse, error) {
	requestDef := GenReqDefForResumeSimultaneousInterpretation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResumeSimultaneousInterpretationResponse), nil
	}
}

// ResumeSimultaneousInterpretationInvoker 开启/关闭同声传译
func (c *MeetingClient) ResumeSimultaneousInterpretationInvoker(request *model.ResumeSimultaneousInterpretationRequest) *ResumeSimultaneousInterpretationInvoker {
	requestDef := GenReqDefForResumeSimultaneousInterpretation()
	return &ResumeSimultaneousInterpretationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RollcallParticipant 点名会场
//
// 该接口用于点名指定与会者。点名会场的效果是除了主持人外，点名与会者为非静音状态，未点名的与会者统一为静音状态。同一时间，只允许一个与会者被点名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) RollcallParticipant(request *model.RollcallParticipantRequest) (*model.RollcallParticipantResponse, error) {
	requestDef := GenReqDefForRollcallParticipant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RollcallParticipantResponse), nil
	}
}

// RollcallParticipantInvoker 点名会场
func (c *MeetingClient) RollcallParticipantInvoker(request *model.RollcallParticipantRequest) *RollcallParticipantInvoker {
	requestDef := GenReqDefForRollcallParticipant()
	return &RollcallParticipantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveLayout 保存多画面布局
//
// 该接口用于保存多画面布局。保存的多画面布局，只能在当前会议使用，会议结束后，保存的多画面布局就会释放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SaveLayout(request *model.SaveLayoutRequest) (*model.SaveLayoutResponse, error) {
	requestDef := GenReqDefForSaveLayout()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveLayoutResponse), nil
	}
}

// SaveLayoutInvoker 保存多画面布局
func (c *MeetingClient) SaveLayoutInvoker(request *model.SaveLayoutRequest) *SaveLayoutInvoker {
	requestDef := GenReqDefForSaveLayout()
	return &SaveLayoutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchAttendanceRecordsOfHisMeeting 查询历史会议的与会者记录
//
// 该接口用于查询指定历史会议的与会者记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchAttendanceRecordsOfHisMeeting(request *model.SearchAttendanceRecordsOfHisMeetingRequest) (*model.SearchAttendanceRecordsOfHisMeetingResponse, error) {
	requestDef := GenReqDefForSearchAttendanceRecordsOfHisMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchAttendanceRecordsOfHisMeetingResponse), nil
	}
}

// SearchAttendanceRecordsOfHisMeetingInvoker 查询历史会议的与会者记录
func (c *MeetingClient) SearchAttendanceRecordsOfHisMeetingInvoker(request *model.SearchAttendanceRecordsOfHisMeetingRequest) *SearchAttendanceRecordsOfHisMeetingInvoker {
	requestDef := GenReqDefForSearchAttendanceRecordsOfHisMeeting()
	return &SearchAttendanceRecordsOfHisMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCorp SP管理员分页搜索企业
//
// SP管理员分页搜索企业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchCorp(request *model.SearchCorpRequest) (*model.SearchCorpResponse, error) {
	requestDef := GenReqDefForSearchCorp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCorpResponse), nil
	}
}

// SearchCorpInvoker SP管理员分页搜索企业
func (c *MeetingClient) SearchCorpInvoker(request *model.SearchCorpRequest) *SearchCorpInvoker {
	requestDef := GenReqDefForSearchCorp()
	return &SearchCorpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCorpAdmins 分页查询企业管理员
//
// 通过该接口分页查询企业管理员。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchCorpAdmins(request *model.SearchCorpAdminsRequest) (*model.SearchCorpAdminsResponse, error) {
	requestDef := GenReqDefForSearchCorpAdmins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCorpAdminsResponse), nil
	}
}

// SearchCorpAdminsInvoker 分页查询企业管理员
func (c *MeetingClient) SearchCorpAdminsInvoker(request *model.SearchCorpAdminsRequest) *SearchCorpAdminsInvoker {
	requestDef := GenReqDefForSearchCorpAdmins()
	return &SearchCorpAdminsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCorpDir 查询企业通讯录
//
// 企业用户（含管理员）通过该接口查询该企业的通讯录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchCorpDir(request *model.SearchCorpDirRequest) (*model.SearchCorpDirResponse, error) {
	requestDef := GenReqDefForSearchCorpDir()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCorpDirResponse), nil
	}
}

// SearchCorpDirInvoker 查询企业通讯录
func (c *MeetingClient) SearchCorpDirInvoker(request *model.SearchCorpDirRequest) *SearchCorpDirInvoker {
	requestDef := GenReqDefForSearchCorpDir()
	return &SearchCorpDirInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCorpExternalDir 查询企业外部联系人
//
// 企业用户（含管理员）通过该接口查询该企业的外部联系人或者个人外部联系人。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchCorpExternalDir(request *model.SearchCorpExternalDirRequest) (*model.SearchCorpExternalDirResponse, error) {
	requestDef := GenReqDefForSearchCorpExternalDir()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCorpExternalDirResponse), nil
	}
}

// SearchCorpExternalDirInvoker 查询企业外部联系人
func (c *MeetingClient) SearchCorpExternalDirInvoker(request *model.SearchCorpExternalDirRequest) *SearchCorpExternalDirInvoker {
	requestDef := GenReqDefForSearchCorpExternalDir()
	return &SearchCorpExternalDirInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCorpResources 企业管理员分页查询企业资源订单列表
//
// 企业管理员根据条件查询企业资源订单列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchCorpResources(request *model.SearchCorpResourcesRequest) (*model.SearchCorpResourcesResponse, error) {
	requestDef := GenReqDefForSearchCorpResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCorpResourcesResponse), nil
	}
}

// SearchCorpResourcesInvoker 企业管理员分页查询企业资源订单列表
func (c *MeetingClient) SearchCorpResourcesInvoker(request *model.SearchCorpResourcesRequest) *SearchCorpResourcesInvoker {
	requestDef := GenReqDefForSearchCorpResources()
	return &SearchCorpResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCorpVmr 企业管理员分页查询企业云会议室
//
// 企业管理员通过该接口分页查询企业的云会议室。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchCorpVmr(request *model.SearchCorpVmrRequest) (*model.SearchCorpVmrResponse, error) {
	requestDef := GenReqDefForSearchCorpVmr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCorpVmrResponse), nil
	}
}

// SearchCorpVmrInvoker 企业管理员分页查询企业云会议室
func (c *MeetingClient) SearchCorpVmrInvoker(request *model.SearchCorpVmrRequest) *SearchCorpVmrInvoker {
	requestDef := GenReqDefForSearchCorpVmr()
	return &SearchCorpVmrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCtlRecordsOfHisMeeting 查询历史会议的会控记录
//
// 该接口用于查询指定历史会议的会控记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchCtlRecordsOfHisMeeting(request *model.SearchCtlRecordsOfHisMeetingRequest) (*model.SearchCtlRecordsOfHisMeetingResponse, error) {
	requestDef := GenReqDefForSearchCtlRecordsOfHisMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCtlRecordsOfHisMeetingResponse), nil
	}
}

// SearchCtlRecordsOfHisMeetingInvoker 查询历史会议的会控记录
func (c *MeetingClient) SearchCtlRecordsOfHisMeetingInvoker(request *model.SearchCtlRecordsOfHisMeetingRequest) *SearchCtlRecordsOfHisMeetingInvoker {
	requestDef := GenReqDefForSearchCtlRecordsOfHisMeeting()
	return &SearchCtlRecordsOfHisMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchDepartmentByName 按名称查询所有的部门
//
// 企业管理员通过该接口按名称查询所有的部门。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchDepartmentByName(request *model.SearchDepartmentByNameRequest) (*model.SearchDepartmentByNameResponse, error) {
	requestDef := GenReqDefForSearchDepartmentByName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchDepartmentByNameResponse), nil
	}
}

// SearchDepartmentByNameInvoker 按名称查询所有的部门
func (c *MeetingClient) SearchDepartmentByNameInvoker(request *model.SearchDepartmentByNameRequest) *SearchDepartmentByNameInvoker {
	requestDef := GenReqDefForSearchDepartmentByName()
	return &SearchDepartmentByNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchDevices 分页查询终端
//
// 企业管理员通过该接口分页查询专业会议终端信息。
// &gt; 如果需要查询Ideahub、SmartRooms、智慧屏TV请使用[[分页查询用户](https://support.huaweicloud.com/api-meeting/meeting_21_0071.html)](tag:hws)[[分页查询用户](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0071.html)](tag:hk)接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchDevices(request *model.SearchDevicesRequest) (*model.SearchDevicesResponse, error) {
	requestDef := GenReqDefForSearchDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchDevicesResponse), nil
	}
}

// SearchDevicesInvoker 分页查询终端
func (c *MeetingClient) SearchDevicesInvoker(request *model.SearchDevicesRequest) *SearchDevicesInvoker {
	requestDef := GenReqDefForSearchDevices()
	return &SearchDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchHisMeetings 查询历史会议列表
//
// 该接口用于查询已经结束的会议。管理员可以查询本企业内所有的历史会议，普通用户仅能查询自己创建或者被邀请的历史会议。不带查询参数时，默认查询权限范围内的历史会议。
// &gt; * 普通用户如果只是通过会议ID或者会议链接接入会议，不是预定者会前邀请或者会中主持人邀请的，则历史会议中无法查到
// &gt; * 如果同一个会议召开并结束多次，则会产生多条历史会议（会议ID相同，会议UUID不同）
// &gt; * 历史会议记录默认保留6个月，最长保留12个月。保留时间管理员可在“会议设置”的“历史会议留存时间”中修改
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchHisMeetings(request *model.SearchHisMeetingsRequest) (*model.SearchHisMeetingsResponse, error) {
	requestDef := GenReqDefForSearchHisMeetings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchHisMeetingsResponse), nil
	}
}

// SearchHisMeetingsInvoker 查询历史会议列表
func (c *MeetingClient) SearchHisMeetingsInvoker(request *model.SearchHisMeetingsRequest) *SearchHisMeetingsInvoker {
	requestDef := GenReqDefForSearchHisMeetings()
	return &SearchHisMeetingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchMaterials 分页查询信息窗素材
//
// 分页查询信息窗素材。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchMaterials(request *model.SearchMaterialsRequest) (*model.SearchMaterialsResponse, error) {
	requestDef := GenReqDefForSearchMaterials()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchMaterialsResponse), nil
	}
}

// SearchMaterialsInvoker 分页查询信息窗素材
func (c *MeetingClient) SearchMaterialsInvoker(request *model.SearchMaterialsRequest) *SearchMaterialsInvoker {
	requestDef := GenReqDefForSearchMaterials()
	return &SearchMaterialsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchMeetingFileList 查询会议纪要列表
//
// 用户查询自己的会议纪要列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchMeetingFileList(request *model.SearchMeetingFileListRequest) (*model.SearchMeetingFileListResponse, error) {
	requestDef := GenReqDefForSearchMeetingFileList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchMeetingFileListResponse), nil
	}
}

// SearchMeetingFileListInvoker 查询会议纪要列表
func (c *MeetingClient) SearchMeetingFileListInvoker(request *model.SearchMeetingFileListRequest) *SearchMeetingFileListInvoker {
	requestDef := GenReqDefForSearchMeetingFileList()
	return &SearchMeetingFileListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchMeetings 查询会议列表
//
// 该接口用于查询尚未结束的会议。
// * 管理员可以查询管理权限域内所有的会议，普通用户仅能查询自己创建或者需要参加的会议。不带查询参数时，默认查询权限范围内正在召开或还未召开的会议。
// * 只能查询尚未结束的会议（既正在召开的会议和已预约还未召开的会议）。如果需要查询历史会议列表，请参考[[查询历史会议列表](https://support.huaweicloud.com/api-meeting/meeting_21_0051.html)](tag:hws)[[查询历史会议列表](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0051.html)](tag:hk)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchMeetings(request *model.SearchMeetingsRequest) (*model.SearchMeetingsResponse, error) {
	requestDef := GenReqDefForSearchMeetings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchMeetingsResponse), nil
	}
}

// SearchMeetingsInvoker 查询会议列表
func (c *MeetingClient) SearchMeetingsInvoker(request *model.SearchMeetingsRequest) *SearchMeetingsInvoker {
	requestDef := GenReqDefForSearchMeetings()
	return &SearchMeetingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchMemberVmr 普通用户分页查询云会议室及个人会议ID
//
// 企业用户通过该接口查询个人已分配的云会议室及个人会议ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchMemberVmr(request *model.SearchMemberVmrRequest) (*model.SearchMemberVmrResponse, error) {
	requestDef := GenReqDefForSearchMemberVmr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchMemberVmrResponse), nil
	}
}

// SearchMemberVmrInvoker 普通用户分页查询云会议室及个人会议ID
func (c *MeetingClient) SearchMemberVmrInvoker(request *model.SearchMemberVmrRequest) *SearchMemberVmrInvoker {
	requestDef := GenReqDefForSearchMemberVmr()
	return &SearchMemberVmrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchOnlineMeetings 查询在线会议列表
//
// 该接口用于查询正在召开的会议列表。管理员可以查询本企业内所有在线会议，普通用户仅能查询当前自己帐号创建或者需要参加的在线会议。不带查询参数时，默认查询权限范围内的在线会议，按开始时间升序排列。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchOnlineMeetings(request *model.SearchOnlineMeetingsRequest) (*model.SearchOnlineMeetingsResponse, error) {
	requestDef := GenReqDefForSearchOnlineMeetings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchOnlineMeetingsResponse), nil
	}
}

// SearchOnlineMeetingsInvoker 查询在线会议列表
func (c *MeetingClient) SearchOnlineMeetingsInvoker(request *model.SearchOnlineMeetingsRequest) *SearchOnlineMeetingsInvoker {
	requestDef := GenReqDefForSearchOnlineMeetings()
	return &SearchOnlineMeetingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchPrograms 查询信息窗节目
//
// 获取信息窗节目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchPrograms(request *model.SearchProgramsRequest) (*model.SearchProgramsResponse, error) {
	requestDef := GenReqDefForSearchPrograms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchProgramsResponse), nil
	}
}

// SearchProgramsInvoker 查询信息窗节目
func (c *MeetingClient) SearchProgramsInvoker(request *model.SearchProgramsRequest) *SearchProgramsInvoker {
	requestDef := GenReqDefForSearchPrograms()
	return &SearchProgramsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchPublications 查询信息窗发布
//
// 获取信息窗发布。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchPublications(request *model.SearchPublicationsRequest) (*model.SearchPublicationsResponse, error) {
	requestDef := GenReqDefForSearchPublications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchPublicationsResponse), nil
	}
}

// SearchPublicationsInvoker 查询信息窗发布
func (c *MeetingClient) SearchPublicationsInvoker(request *model.SearchPublicationsRequest) *SearchPublicationsInvoker {
	requestDef := GenReqDefForSearchPublications()
	return &SearchPublicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchRecordings 查询录制列表
//
// 该接口用于查询会议录制列表。管理员可以查询本企业内所有的录制，普通用户仅能查询自己创建的会议的录制。不带查询参数时，默认查询权限范围内的录制。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchRecordings(request *model.SearchRecordingsRequest) (*model.SearchRecordingsResponse, error) {
	requestDef := GenReqDefForSearchRecordings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchRecordingsResponse), nil
	}
}

// SearchRecordingsInvoker 查询录制列表
func (c *MeetingClient) SearchRecordingsInvoker(request *model.SearchRecordingsRequest) *SearchRecordingsInvoker {
	requestDef := GenReqDefForSearchRecordings()
	return &SearchRecordingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchResource SP管理员根据分页查询企业资源
//
// SP根据条件查询企业的资源项。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchResource(request *model.SearchResourceRequest) (*model.SearchResourceResponse, error) {
	requestDef := GenReqDefForSearchResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchResourceResponse), nil
	}
}

// SearchResourceInvoker SP管理员根据分页查询企业资源
func (c *MeetingClient) SearchResourceInvoker(request *model.SearchResourceRequest) *SearchResourceInvoker {
	requestDef := GenReqDefForSearchResource()
	return &SearchResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchResourceOpRecord SP管理员分页查询企业资源操作记录
//
// SP根据根据条件查询企业的资源操作记录，支持根据resourceId模糊搜索。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchResourceOpRecord(request *model.SearchResourceOpRecordRequest) (*model.SearchResourceOpRecordResponse, error) {
	requestDef := GenReqDefForSearchResourceOpRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchResourceOpRecordResponse), nil
	}
}

// SearchResourceOpRecordInvoker SP管理员分页查询企业资源操作记录
func (c *MeetingClient) SearchResourceOpRecordInvoker(request *model.SearchResourceOpRecordRequest) *SearchResourceOpRecordInvoker {
	requestDef := GenReqDefForSearchResourceOpRecord()
	return &SearchResourceOpRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchUsers 分页查询用户
//
// 企业管理员通过该接口分页查询企业用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchUsers(request *model.SearchUsersRequest) (*model.SearchUsersResponse, error) {
	requestDef := GenReqDefForSearchUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchUsersResponse), nil
	}
}

// SearchUsersInvoker 分页查询用户
func (c *MeetingClient) SearchUsersInvoker(request *model.SearchUsersRequest) *SearchUsersInvoker {
	requestDef := GenReqDefForSearchUsers()
	return &SearchUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchVisionActiveCode 企业管理员分页查询激活码
//
// 企业管理员分页查询激活码，支持激活码、终端名称模糊查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchVisionActiveCode(request *model.SearchVisionActiveCodeRequest) (*model.SearchVisionActiveCodeResponse, error) {
	requestDef := GenReqDefForSearchVisionActiveCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchVisionActiveCodeResponse), nil
	}
}

// SearchVisionActiveCodeInvoker 企业管理员分页查询激活码
func (c *MeetingClient) SearchVisionActiveCodeInvoker(request *model.SearchVisionActiveCodeRequest) *SearchVisionActiveCodeInvoker {
	requestDef := GenReqDefForSearchVisionActiveCode()
	return &SearchVisionActiveCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendSlideVerifyCode 发送滑块验证码
//
// 该接口提供发送滑块验证码。服务器收到请求，返回抠图以及抠图后的原图等结果。需要在前台界面显示出抠图以及抠图后的原图，用户通过滑块操作来匹配图形。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SendSlideVerifyCode(request *model.SendSlideVerifyCodeRequest) (*model.SendSlideVerifyCodeResponse, error) {
	requestDef := GenReqDefForSendSlideVerifyCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendSlideVerifyCodeResponse), nil
	}
}

// SendSlideVerifyCodeInvoker 发送滑块验证码
func (c *MeetingClient) SendSlideVerifyCodeInvoker(request *model.SendSlideVerifyCodeRequest) *SendSlideVerifyCodeInvoker {
	requestDef := GenReqDefForSendSlideVerifyCode()
	return &SendSlideVerifyCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendVeriCodeForChangePwd 发送验证码
//
// 该接口提供发送验证码，服务器收到请求，发送验证码到邮箱或者短信并返回结果。用户在前台界面通过滑块验证后，再进行发送验证码操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SendVeriCodeForChangePwd(request *model.SendVeriCodeForChangePwdRequest) (*model.SendVeriCodeForChangePwdResponse, error) {
	requestDef := GenReqDefForSendVeriCodeForChangePwd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendVeriCodeForChangePwdResponse), nil
	}
}

// SendVeriCodeForChangePwdInvoker 发送验证码
func (c *MeetingClient) SendVeriCodeForChangePwdInvoker(request *model.SendVeriCodeForChangePwdRequest) *SendVeriCodeForChangePwdInvoker {
	requestDef := GenReqDefForSendVeriCodeForChangePwd()
	return &SendVeriCodeForChangePwdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendVeriCodeForUpdateUserInfo 获取验证码
//
// 修改用户手机或邮箱时，需要获取验证码。企业用户通过该接口获取验证码，系统会向用户的手机或邮箱发送，验证码1分钟内有效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SendVeriCodeForUpdateUserInfo(request *model.SendVeriCodeForUpdateUserInfoRequest) (*model.SendVeriCodeForUpdateUserInfoResponse, error) {
	requestDef := GenReqDefForSendVeriCodeForUpdateUserInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendVeriCodeForUpdateUserInfoResponse), nil
	}
}

// SendVeriCodeForUpdateUserInfoInvoker 获取验证码
func (c *MeetingClient) SendVeriCodeForUpdateUserInfoInvoker(request *model.SendVeriCodeForUpdateUserInfoRequest) *SendVeriCodeForUpdateUserInfoInvoker {
	requestDef := GenReqDefForSendVeriCodeForUpdateUserInfo()
	return &SendVeriCodeForUpdateUserInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetAttendeeLanChannel 设置普通与会人的语言频道
//
// 主持人通过该接口设置某些普通与会者(包括主持人)加入哪个语言频道。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SetAttendeeLanChannel(request *model.SetAttendeeLanChannelRequest) (*model.SetAttendeeLanChannelResponse, error) {
	requestDef := GenReqDefForSetAttendeeLanChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAttendeeLanChannelResponse), nil
	}
}

// SetAttendeeLanChannelInvoker 设置普通与会人的语言频道
func (c *MeetingClient) SetAttendeeLanChannelInvoker(request *model.SetAttendeeLanChannelRequest) *SetAttendeeLanChannelInvoker {
	requestDef := GenReqDefForSetAttendeeLanChannel()
	return &SetAttendeeLanChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetCohost 申请联席主持人
//
// 该接口用于设置联席主持人或释放联席主持人。只能将来宾设置为联席主持人。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SetCohost(request *model.SetCohostRequest) (*model.SetCohostResponse, error) {
	requestDef := GenReqDefForSetCohost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetCohostResponse), nil
	}
}

// SetCohostInvoker 申请联席主持人
func (c *MeetingClient) SetCohostInvoker(request *model.SetCohostRequest) *SetCohostInvoker {
	requestDef := GenReqDefForSetCohost()
	return &SetCohostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetCustomMultiPicture 设置自定义多画面
//
// 该接口用于设置会中多画面。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SetCustomMultiPicture(request *model.SetCustomMultiPictureRequest) (*model.SetCustomMultiPictureResponse, error) {
	requestDef := GenReqDefForSetCustomMultiPicture()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetCustomMultiPictureResponse), nil
	}
}

// SetCustomMultiPictureInvoker 设置自定义多画面
func (c *MeetingClient) SetCustomMultiPictureInvoker(request *model.SetCustomMultiPictureRequest) *SetCustomMultiPictureInvoker {
	requestDef := GenReqDefForSetCustomMultiPicture()
	return &SetCustomMultiPictureInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetHostView 主持人选看
//
// 该接口用于主持人轮询、主持人选看多画面、主持人选看会场操作。只适用于专业会议终端（如TE系列等）为主持人的场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SetHostView(request *model.SetHostViewRequest) (*model.SetHostViewResponse, error) {
	requestDef := GenReqDefForSetHostView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetHostViewResponse), nil
	}
}

// SetHostViewInvoker 主持人选看
func (c *MeetingClient) SetHostViewInvoker(request *model.SetHostViewRequest) *SetHostViewInvoker {
	requestDef := GenReqDefForSetHostView()
	return &SetHostViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetInterpreterGroup 设置传译组
//
// 主持人通过该接口设置传译组，每个传译组支持两种语言，传译组内支持多个传译员。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SetInterpreterGroup(request *model.SetInterpreterGroupRequest) (*model.SetInterpreterGroupResponse, error) {
	requestDef := GenReqDefForSetInterpreterGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetInterpreterGroupResponse), nil
	}
}

// SetInterpreterGroupInvoker 设置传译组
func (c *MeetingClient) SetInterpreterGroupInvoker(request *model.SetInterpreterGroupRequest) *SetInterpreterGroupInvoker {
	requestDef := GenReqDefForSetInterpreterGroup()
	return &SetInterpreterGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetMultiPicture 设置多画面
//
// 设置会议多画面。该接口废弃不用，请使用“[[设置自定义多画面](https://support.huaweicloud.com/api-meeting/meeting_21_0418.html)](tag:hws)[[设置自定义多画面](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0418.html)](tag:hk)”接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SetMultiPicture(request *model.SetMultiPictureRequest) (*model.SetMultiPictureResponse, error) {
	requestDef := GenReqDefForSetMultiPicture()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetMultiPictureResponse), nil
	}
}

// SetMultiPictureInvoker 设置多画面
func (c *MeetingClient) SetMultiPictureInvoker(request *model.SetMultiPictureRequest) *SetMultiPictureInvoker {
	requestDef := GenReqDefForSetMultiPicture()
	return &SetMultiPictureInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetParticipantView 会场选看
//
// 该接口用于专业会议终端（如TE系列等）选看其他与会者。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SetParticipantView(request *model.SetParticipantViewRequest) (*model.SetParticipantViewResponse, error) {
	requestDef := GenReqDefForSetParticipantView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetParticipantViewResponse), nil
	}
}

// SetParticipantViewInvoker 会场选看
func (c *MeetingClient) SetParticipantViewInvoker(request *model.SetParticipantViewRequest) *SetParticipantViewInvoker {
	requestDef := GenReqDefForSetParticipantView()
	return &SetParticipantViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetProfileImage 用户设置头像
//
// 用户设置头像
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SetProfileImage(request *model.SetProfileImageRequest) (*model.SetProfileImageResponse, error) {
	requestDef := GenReqDefForSetProfileImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetProfileImageResponse), nil
	}
}

// SetProfileImageInvoker 用户设置头像
func (c *MeetingClient) SetProfileImageInvoker(request *model.SetProfileImageRequest) *SetProfileImageInvoker {
	requestDef := GenReqDefForSetProfileImage()
	return &SetProfileImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetRole 申请主持人
//
// 该接口用于设置主持人或释放主持人。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SetRole(request *model.SetRoleRequest) (*model.SetRoleResponse, error) {
	requestDef := GenReqDefForSetRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRoleResponse), nil
	}
}

// SetRoleInvoker 申请主持人
func (c *MeetingClient) SetRoleInvoker(request *model.SetRoleRequest) *SetRoleInvoker {
	requestDef := GenReqDefForSetRole()
	return &SetRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetSsoConfig 设置SSO登录配置
//
// 该接口用于设置SSO登录的鉴权配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SetSsoConfig(request *model.SetSsoConfigRequest) (*model.SetSsoConfigResponse, error) {
	requestDef := GenReqDefForSetSsoConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetSsoConfigResponse), nil
	}
}

// SetSsoConfigInvoker 设置SSO登录配置
func (c *MeetingClient) SetSsoConfigInvoker(request *model.SetSsoConfigRequest) *SetSsoConfigInvoker {
	requestDef := GenReqDefForSetSsoConfig()
	return &SetSsoConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetUserProfileImage 企业管理员设置企业成员头像
//
// 为企业内的用户设置头像（只允许管理员调用）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SetUserProfileImage(request *model.SetUserProfileImageRequest) (*model.SetUserProfileImageResponse, error) {
	requestDef := GenReqDefForSetUserProfileImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetUserProfileImageResponse), nil
	}
}

// SetUserProfileImageInvoker 企业管理员设置企业成员头像
func (c *MeetingClient) SetUserProfileImageInvoker(request *model.SetUserProfileImageRequest) *SetUserProfileImageInvoker {
	requestDef := GenReqDefForSetUserProfileImage()
	return &SetUserProfileImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetWebHookConfig 设置事件推送
//
// 该接口用于管理员设置企业级会议事件订阅配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SetWebHookConfig(request *model.SetWebHookConfigRequest) (*model.SetWebHookConfigResponse, error) {
	requestDef := GenReqDefForSetWebHookConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetWebHookConfigResponse), nil
	}
}

// SetWebHookConfigInvoker 设置事件推送
func (c *MeetingClient) SetWebHookConfigInvoker(request *model.SetWebHookConfigRequest) *SetWebHookConfigInvoker {
	requestDef := GenReqDefForSetWebHookConfig()
	return &SetWebHookConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfOrg SP管理员查询会议归属企业
//
// SP管理员根据会议ID查询该会议归属的企业ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowConfOrg(request *model.ShowConfOrgRequest) (*model.ShowConfOrgResponse, error) {
	requestDef := GenReqDefForShowConfOrg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfOrgResponse), nil
	}
}

// ShowConfOrgInvoker SP管理员查询会议归属企业
func (c *MeetingClient) ShowConfOrgInvoker(request *model.ShowConfOrgRequest) *ShowConfOrgInvoker {
	requestDef := GenReqDefForShowConfOrg()
	return &ShowConfOrgInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCorp SP管理员查询企业
//
// 获取企业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowCorp(request *model.ShowCorpRequest) (*model.ShowCorpResponse, error) {
	requestDef := GenReqDefForShowCorp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCorpResponse), nil
	}
}

// ShowCorpInvoker SP管理员查询企业
func (c *MeetingClient) ShowCorpInvoker(request *model.ShowCorpRequest) *ShowCorpInvoker {
	requestDef := GenReqDefForShowCorp()
	return &ShowCorpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCorpAdmin 查询企业管理员
//
// 通过该接口查询企业管理员。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowCorpAdmin(request *model.ShowCorpAdminRequest) (*model.ShowCorpAdminResponse, error) {
	requestDef := GenReqDefForShowCorpAdmin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCorpAdminResponse), nil
	}
}

// ShowCorpAdminInvoker 查询企业管理员
func (c *MeetingClient) ShowCorpAdminInvoker(request *model.ShowCorpAdminRequest) *ShowCorpAdminInvoker {
	requestDef := GenReqDefForShowCorpAdmin()
	return &ShowCorpAdminInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCorpBasicInfo 企业管理员查询企业注册信息
//
// 企业管理员通过该接口查询企业注册信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowCorpBasicInfo(request *model.ShowCorpBasicInfoRequest) (*model.ShowCorpBasicInfoResponse, error) {
	requestDef := GenReqDefForShowCorpBasicInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCorpBasicInfoResponse), nil
	}
}

// ShowCorpBasicInfoInvoker 企业管理员查询企业注册信息
func (c *MeetingClient) ShowCorpBasicInfoInvoker(request *model.ShowCorpBasicInfoRequest) *ShowCorpBasicInfoInvoker {
	requestDef := GenReqDefForShowCorpBasicInfo()
	return &ShowCorpBasicInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCorpResource 企业管理员查询企业内资源及业务权限
//
// 企业管理员通过该接口查询企业内资源及业务权限，包括查询已使用的资源情况。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowCorpResource(request *model.ShowCorpResourceRequest) (*model.ShowCorpResourceResponse, error) {
	requestDef := GenReqDefForShowCorpResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCorpResourceResponse), nil
	}
}

// ShowCorpResourceInvoker 企业管理员查询企业内资源及业务权限
func (c *MeetingClient) ShowCorpResourceInvoker(request *model.ShowCorpResourceRequest) *ShowCorpResourceInvoker {
	requestDef := GenReqDefForShowCorpResource()
	return &ShowCorpResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDepartment 通过部门编码查询部门信息
//
// 通过部门编码查询部门信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowDepartment(request *model.ShowDepartmentRequest) (*model.ShowDepartmentResponse, error) {
	requestDef := GenReqDefForShowDepartment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDepartmentResponse), nil
	}
}

// ShowDepartmentInvoker 通过部门编码查询部门信息
func (c *MeetingClient) ShowDepartmentInvoker(request *model.ShowDepartmentRequest) *ShowDepartmentInvoker {
	requestDef := GenReqDefForShowDepartment()
	return &ShowDepartmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeptAndChildDept 查询部门及其一级子部门列表
//
// 企业管理员通过该接口查询部门及其一级子部门列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowDeptAndChildDept(request *model.ShowDeptAndChildDeptRequest) (*model.ShowDeptAndChildDeptResponse, error) {
	requestDef := GenReqDefForShowDeptAndChildDept()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeptAndChildDeptResponse), nil
	}
}

// ShowDeptAndChildDeptInvoker 查询部门及其一级子部门列表
func (c *MeetingClient) ShowDeptAndChildDeptInvoker(request *model.ShowDeptAndChildDeptRequest) *ShowDeptAndChildDeptInvoker {
	requestDef := GenReqDefForShowDeptAndChildDept()
	return &ShowDeptAndChildDeptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeviceDetail 查询终端详情
//
// 企业管理员通过该接口查询专业会议终端详情。
// &gt; 如果需要查询Ideahub、SmartRooms、智慧屏TV详情请使用[[查询用户详情](https://support.huaweicloud.com/api-meeting/meeting_21_0069.html)](tag:hws)[[查询用户详情](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0069.html)](tag:hk)接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowDeviceDetail(request *model.ShowDeviceDetailRequest) (*model.ShowDeviceDetailResponse, error) {
	requestDef := GenReqDefForShowDeviceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceDetailResponse), nil
	}
}

// ShowDeviceDetailInvoker 查询终端详情
func (c *MeetingClient) ShowDeviceDetailInvoker(request *model.ShowDeviceDetailRequest) *ShowDeviceDetailInvoker {
	requestDef := GenReqDefForShowDeviceDetail()
	return &ShowDeviceDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeviceStatus 查询设备状态
//
// 调用本接口可以查询硬终端的状态。
// 硬终端与发起查询请求的帐号需在同一企业下，否则会鉴权失败。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowDeviceStatus(request *model.ShowDeviceStatusRequest) (*model.ShowDeviceStatusResponse, error) {
	requestDef := GenReqDefForShowDeviceStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceStatusResponse), nil
	}
}

// ShowDeviceStatusInvoker 查询设备状态
func (c *MeetingClient) ShowDeviceStatusInvoker(request *model.ShowDeviceStatusRequest) *ShowDeviceStatusInvoker {
	requestDef := GenReqDefForShowDeviceStatus()
	return &ShowDeviceStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeviceTypes 获取所有终端类型
//
// 企业管理员通过该接口获取所有的专业会议终端类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowDeviceTypes(request *model.ShowDeviceTypesRequest) (*model.ShowDeviceTypesResponse, error) {
	requestDef := GenReqDefForShowDeviceTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceTypesResponse), nil
	}
}

// ShowDeviceTypesInvoker 获取所有终端类型
func (c *MeetingClient) ShowDeviceTypesInvoker(request *model.ShowDeviceTypesRequest) *ShowDeviceTypesInvoker {
	requestDef := GenReqDefForShowDeviceTypes()
	return &ShowDeviceTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHisMeetingDetail 查询历史会议详情
//
// 该接口用户查询指定历史会议的详情。管理员可以查询本企业内所有的历史会议详情，普通用户仅能查询自己创建或者被邀请的历史会议详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowHisMeetingDetail(request *model.ShowHisMeetingDetailRequest) (*model.ShowHisMeetingDetailResponse, error) {
	requestDef := GenReqDefForShowHisMeetingDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHisMeetingDetailResponse), nil
	}
}

// ShowHisMeetingDetailInvoker 查询历史会议详情
func (c *MeetingClient) ShowHisMeetingDetailInvoker(request *model.ShowHisMeetingDetailRequest) *ShowHisMeetingDetailInvoker {
	requestDef := GenReqDefForShowHisMeetingDetail()
	return &ShowHisMeetingDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLayout 查询多画面布局
//
// 该接口用于查询当前会议已保存的多画面布局。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowLayout(request *model.ShowLayoutRequest) (*model.ShowLayoutResponse, error) {
	requestDef := GenReqDefForShowLayout()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLayoutResponse), nil
	}
}

// ShowLayoutInvoker 查询多画面布局
func (c *MeetingClient) ShowLayoutInvoker(request *model.ShowLayoutRequest) *ShowLayoutInvoker {
	requestDef := GenReqDefForShowLayout()
	return &ShowLayoutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMeetingDetail 查询会议详情
//
// 查询偏移量
// * 管理员可以查询管理权限域内所有会议的详情，普通用户仅能查询自己创建或者需要参加的会议详情。
// * 只能查询尚未结束的会议（既正在召开的会议和已预约还未召开的会议）。如果需要查询历史会议列详情，请参考[[查询历史会议详情](https://support.huaweicloud.com/api-meeting/meeting_21_0052.html)](tag:hws)[[查询历史会议详情](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0052.html)](tag:hk)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowMeetingDetail(request *model.ShowMeetingDetailRequest) (*model.ShowMeetingDetailResponse, error) {
	requestDef := GenReqDefForShowMeetingDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMeetingDetailResponse), nil
	}
}

// ShowMeetingDetailInvoker 查询会议详情
func (c *MeetingClient) ShowMeetingDetailInvoker(request *model.ShowMeetingDetailRequest) *ShowMeetingDetailInvoker {
	requestDef := GenReqDefForShowMeetingDetail()
	return &ShowMeetingDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMeetingFile 查询会议纪要详情
//
// 用户查询单个会议纪要详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowMeetingFile(request *model.ShowMeetingFileRequest) (*model.ShowMeetingFileResponse, error) {
	requestDef := GenReqDefForShowMeetingFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMeetingFileResponse), nil
	}
}

// ShowMeetingFileInvoker 查询会议纪要详情
func (c *MeetingClient) ShowMeetingFileInvoker(request *model.ShowMeetingFileRequest) *ShowMeetingFileInvoker {
	requestDef := GenReqDefForShowMeetingFile()
	return &ShowMeetingFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMeetingFileList 打开会议纪要文件列表
//
// 用户使用手机扫码后，手机端请求服务端,让服务端通知指定IdeaHub打开指定用户的会议纪要文件列表。二维码内容 ：cloudlink://cloudlink.huawei.com/h5page?action&#x3D;OPEN_MEETING_FILE_LIST&amp;key1&#x3D;value1&amp;key2&#x3D;value2 。key/value的个数可能变化，终端解析后，在发起后续请求时，将所有key/value存为map，作为入参即可。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowMeetingFileList(request *model.ShowMeetingFileListRequest) (*model.ShowMeetingFileListResponse, error) {
	requestDef := GenReqDefForShowMeetingFileList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMeetingFileListResponse), nil
	}
}

// ShowMeetingFileListInvoker 打开会议纪要文件列表
func (c *MeetingClient) ShowMeetingFileListInvoker(request *model.ShowMeetingFileListRequest) *ShowMeetingFileListInvoker {
	requestDef := GenReqDefForShowMeetingFileList()
	return &ShowMeetingFileListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMyInfo 用户查询自己的信息
//
// 企业用户通过该接口查询自己的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowMyInfo(request *model.ShowMyInfoRequest) (*model.ShowMyInfoResponse, error) {
	requestDef := GenReqDefForShowMyInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMyInfoResponse), nil
	}
}

// ShowMyInfoInvoker 用户查询自己的信息
func (c *MeetingClient) ShowMyInfoInvoker(request *model.ShowMyInfoRequest) *ShowMyInfoInvoker {
	requestDef := GenReqDefForShowMyInfo()
	return &ShowMyInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOnlineMeetingDetail 查询在线会议详情
//
// 该接口用于查询正在召开的会议详情。管理员可以查询本企业内所有的在线会议详情，普通用户仅能查询自己帐号创建或者需要参加的在线会议详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowOnlineMeetingDetail(request *model.ShowOnlineMeetingDetailRequest) (*model.ShowOnlineMeetingDetailResponse, error) {
	requestDef := GenReqDefForShowOnlineMeetingDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOnlineMeetingDetailResponse), nil
	}
}

// ShowOnlineMeetingDetailInvoker 查询在线会议详情
func (c *MeetingClient) ShowOnlineMeetingDetailInvoker(request *model.ShowOnlineMeetingDetailRequest) *ShowOnlineMeetingDetailInvoker {
	requestDef := GenReqDefForShowOnlineMeetingDetail()
	return &ShowOnlineMeetingDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrgRes 企业管理员查询企业资源使用信息
//
// 企业管理员查询所属企业的资源使用信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowOrgRes(request *model.ShowOrgResRequest) (*model.ShowOrgResResponse, error) {
	requestDef := GenReqDefForShowOrgRes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrgResResponse), nil
	}
}

// ShowOrgResInvoker 企业管理员查询企业资源使用信息
func (c *MeetingClient) ShowOrgResInvoker(request *model.ShowOrgResRequest) *ShowOrgResInvoker {
	requestDef := GenReqDefForShowOrgRes()
	return &ShowOrgResInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProgram 根据ID查询节目详情
//
// 根据ID获取节目详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowProgram(request *model.ShowProgramRequest) (*model.ShowProgramResponse, error) {
	requestDef := GenReqDefForShowProgram()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProgramResponse), nil
	}
}

// ShowProgramInvoker 根据ID查询节目详情
func (c *MeetingClient) ShowProgramInvoker(request *model.ShowProgramRequest) *ShowProgramInvoker {
	requestDef := GenReqDefForShowProgram()
	return &ShowProgramInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublication 根据ID查询信息窗发布详情
//
// 根据ID获取发布详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowPublication(request *model.ShowPublicationRequest) (*model.ShowPublicationResponse, error) {
	requestDef := GenReqDefForShowPublication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicationResponse), nil
	}
}

// ShowPublicationInvoker 根据ID查询信息窗发布详情
func (c *MeetingClient) ShowPublicationInvoker(request *model.ShowPublicationRequest) *ShowPublicationInvoker {
	requestDef := GenReqDefForShowPublication()
	return &ShowPublicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRealTimeInfoOfMeeting 查询会议实时信息
//
// 该接口用于查询正在召开的会议实时信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowRealTimeInfoOfMeeting(request *model.ShowRealTimeInfoOfMeetingRequest) (*model.ShowRealTimeInfoOfMeetingResponse, error) {
	requestDef := GenReqDefForShowRealTimeInfoOfMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRealTimeInfoOfMeetingResponse), nil
	}
}

// ShowRealTimeInfoOfMeetingInvoker 查询会议实时信息
func (c *MeetingClient) ShowRealTimeInfoOfMeetingInvoker(request *model.ShowRealTimeInfoOfMeetingRequest) *ShowRealTimeInfoOfMeetingInvoker {
	requestDef := GenReqDefForShowRealTimeInfoOfMeeting()
	return &ShowRealTimeInfoOfMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecordInfo 查询单会议录制文件信息
//
// 查询单会议录制文件信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowRecordInfo(request *model.ShowRecordInfoRequest) (*model.ShowRecordInfoResponse, error) {
	requestDef := GenReqDefForShowRecordInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordInfoResponse), nil
	}
}

// ShowRecordInfoInvoker 查询单会议录制文件信息
func (c *MeetingClient) ShowRecordInfoInvoker(request *model.ShowRecordInfoRequest) *ShowRecordInfoInvoker {
	requestDef := GenReqDefForShowRecordInfo()
	return &ShowRecordInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecordingDetail 查询录制详情
//
// 改接口用于查询某个会议录制的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowRecordingDetail(request *model.ShowRecordingDetailRequest) (*model.ShowRecordingDetailResponse, error) {
	requestDef := GenReqDefForShowRecordingDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordingDetailResponse), nil
	}
}

// ShowRecordingDetailInvoker 查询录制详情
func (c *MeetingClient) ShowRecordingDetailInvoker(request *model.ShowRecordingDetailRequest) *ShowRecordingDetailInvoker {
	requestDef := GenReqDefForShowRecordingDetail()
	return &ShowRecordingDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecordingFileDownloadUrls 查询录制文件下载链接
//
// 该接口用户查询指定会议录制文件下载链接。
// &gt; * 仅企业管理员权限的帐号才能查询录制文件的下载链接
// &gt; * 这个接口需要在华为云会议后台开通白名单后才能调用。请联系华为销售人员，并提供华为云会议企业ID
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowRecordingFileDownloadUrls(request *model.ShowRecordingFileDownloadUrlsRequest) (*model.ShowRecordingFileDownloadUrlsResponse, error) {
	requestDef := GenReqDefForShowRecordingFileDownloadUrls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordingFileDownloadUrlsResponse), nil
	}
}

// ShowRecordingFileDownloadUrlsInvoker 查询录制文件下载链接
func (c *MeetingClient) ShowRecordingFileDownloadUrlsInvoker(request *model.ShowRecordingFileDownloadUrlsRequest) *ShowRecordingFileDownloadUrlsInvoker {
	requestDef := GenReqDefForShowRecordingFileDownloadUrls()
	return &ShowRecordingFileDownloadUrlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRegionInfoOfMeeting 查询会议所在区域信息
//
// 该接口用于查询会议所在区域的IP和域名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowRegionInfoOfMeeting(request *model.ShowRegionInfoOfMeetingRequest) (*model.ShowRegionInfoOfMeetingResponse, error) {
	requestDef := GenReqDefForShowRegionInfoOfMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRegionInfoOfMeetingResponse), nil
	}
}

// ShowRegionInfoOfMeetingInvoker 查询会议所在区域信息
func (c *MeetingClient) ShowRegionInfoOfMeetingInvoker(request *model.ShowRegionInfoOfMeetingRequest) *ShowRegionInfoOfMeetingInvoker {
	requestDef := GenReqDefForShowRegionInfoOfMeeting()
	return &ShowRegionInfoOfMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRoomSetting 查询网络研讨会高级设置
//
// 该接口用于查询指定网络研讨会的高级设置。管理员可查询企业内的网络研讨会高级设置，非管理员只可查询自己预定的网络研讨会的高级设置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowRoomSetting(request *model.ShowRoomSettingRequest) (*model.ShowRoomSettingResponse, error) {
	requestDef := GenReqDefForShowRoomSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRoomSettingResponse), nil
	}
}

// ShowRoomSettingInvoker 查询网络研讨会高级设置
func (c *MeetingClient) ShowRoomSettingInvoker(request *model.ShowRoomSettingRequest) *ShowRoomSettingInvoker {
	requestDef := GenReqDefForShowRoomSetting()
	return &ShowRoomSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSpRes SP管理员查询SP下资源使用信息
//
// SP管理员查询所属SP的共享资源使用信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowSpRes(request *model.ShowSpResRequest) (*model.ShowSpResResponse, error) {
	requestDef := GenReqDefForShowSpRes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSpResResponse), nil
	}
}

// ShowSpResInvoker SP管理员查询SP下资源使用信息
func (c *MeetingClient) ShowSpResInvoker(request *model.ShowSpResRequest) *ShowSpResInvoker {
	requestDef := GenReqDefForShowSpRes()
	return &ShowSpResInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSpResource SP管理员查询资源信息
//
// SP管理员查询SP的所有资源，包括已使用的资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowSpResource(request *model.ShowSpResourceRequest) (*model.ShowSpResourceResponse, error) {
	requestDef := GenReqDefForShowSpResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSpResourceResponse), nil
	}
}

// ShowSpResourceInvoker SP管理员查询资源信息
func (c *MeetingClient) ShowSpResourceInvoker(request *model.ShowSpResourceRequest) *ShowSpResourceInvoker {
	requestDef := GenReqDefForShowSpResource()
	return &ShowSpResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSsoConfig 查询SSO登录配置
//
// 该接口用于查询SSO登录的鉴权配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowSsoConfig(request *model.ShowSsoConfigRequest) (*model.ShowSsoConfigResponse, error) {
	requestDef := GenReqDefForShowSsoConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSsoConfigResponse), nil
	}
}

// ShowSsoConfigInvoker 查询SSO登录配置
func (c *MeetingClient) ShowSsoConfigInvoker(request *model.ShowSsoConfigRequest) *ShowSsoConfigInvoker {
	requestDef := GenReqDefForShowSsoConfig()
	return &ShowSsoConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserDetail 查询用户详情
//
// 企业管理员通过该接口查询企业用户详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowUserDetail(request *model.ShowUserDetailRequest) (*model.ShowUserDetailResponse, error) {
	requestDef := GenReqDefForShowUserDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserDetailResponse), nil
	}
}

// ShowUserDetailInvoker 查询用户详情
func (c *MeetingClient) ShowUserDetailInvoker(request *model.ShowUserDetailRequest) *ShowUserDetailInvoker {
	requestDef := GenReqDefForShowUserDetail()
	return &ShowUserDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWebHookConfig 查询事件推送
//
// 该接口用于管理员查询企业事件订阅配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowWebHookConfig(request *model.ShowWebHookConfigRequest) (*model.ShowWebHookConfigResponse, error) {
	requestDef := GenReqDefForShowWebHookConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWebHookConfigResponse), nil
	}
}

// ShowWebHookConfigInvoker 查询事件推送
func (c *MeetingClient) ShowWebHookConfigInvoker(request *model.ShowWebHookConfigRequest) *ShowWebHookConfigInvoker {
	requestDef := GenReqDefForShowWebHookConfig()
	return &ShowWebHookConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWebinar 查询网络研讨会详情
//
// 该接口用于查询指定网络研讨会的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowWebinar(request *model.ShowWebinarRequest) (*model.ShowWebinarResponse, error) {
	requestDef := GenReqDefForShowWebinar()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWebinarResponse), nil
	}
}

// ShowWebinarInvoker 查询网络研讨会详情
func (c *MeetingClient) ShowWebinarInvoker(request *model.ShowWebinarRequest) *ShowWebinarInvoker {
	requestDef := GenReqDefForShowWebinar()
	return &ShowWebinarInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartMeeting 激活会议
//
// 该接口用于通过会议ID和会议密码激活会议。所有的会控接口都需要在会议激活后才能调用，可以通过该接口先激活会议。
// &gt; 来宾密码是否可以激活会议取决于会议创建时是否设置了“是否允许来宾启动会议”（allowGuestStartConf&#x3D;true）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) StartMeeting(request *model.StartMeetingRequest) (*model.StartMeetingResponse, error) {
	requestDef := GenReqDefForStartMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartMeetingResponse), nil
	}
}

// StartMeetingInvoker 激活会议
func (c *MeetingClient) StartMeetingInvoker(request *model.StartMeetingRequest) *StartMeetingInvoker {
	requestDef := GenReqDefForStartMeeting()
	return &StartMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopMeeting 结束会议
//
// 该接口用于结束正在召开的会议。
// &gt; * 如果管理员在企业的会议设置中关闭“结束会议保留预约记录”开关，会议结束后会议列表中将删除该会议，与会者不能再次加入该会议。否则会议预约时间到之前，与会者可以再次加入该会议
// &gt; * “结束会议保留预约记录”默认是开的
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) StopMeeting(request *model.StopMeetingRequest) (*model.StopMeetingResponse, error) {
	requestDef := GenReqDefForStopMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopMeetingResponse), nil
	}
}

// StopMeetingInvoker 结束会议
func (c *MeetingClient) StopMeetingInvoker(request *model.StopMeetingRequest) *StopMeetingInvoker {
	requestDef := GenReqDefForStopMeeting()
	return &StopMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchMode 切换视频显示策略
//
// 该接口用于切换会中视频画面显示策略，包括广播多画面，广播单画面，声控多画面。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SwitchMode(request *model.SwitchModeRequest) (*model.SwitchModeResponse, error) {
	requestDef := GenReqDefForSwitchMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchModeResponse), nil
	}
}

// SwitchModeInvoker 切换视频显示策略
func (c *MeetingClient) SwitchModeInvoker(request *model.SwitchModeRequest) *SwitchModeInvoker {
	requestDef := GenReqDefForSwitchMode()
	return &SwitchModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAppId 修改企业应用
//
// 企业默认管理员修改企业应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateAppId(request *model.UpdateAppIdRequest) (*model.UpdateAppIdResponse, error) {
	requestDef := GenReqDefForUpdateAppId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppIdResponse), nil
	}
}

// UpdateAppIdInvoker 修改企业应用
func (c *MeetingClient) UpdateAppIdInvoker(request *model.UpdateAppIdRequest) *UpdateAppIdInvoker {
	requestDef := GenReqDefForUpdateAppId()
	return &UpdateAppIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateContact 修改手机或邮箱
//
// 企业用户通过该接口修改手机或邮箱，需要先获取验证码，验证多次失败会禁止修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateContact(request *model.UpdateContactRequest) (*model.UpdateContactResponse, error) {
	requestDef := GenReqDefForUpdateContact()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateContactResponse), nil
	}
}

// UpdateContactInvoker 修改手机或邮箱
func (c *MeetingClient) UpdateContactInvoker(request *model.UpdateContactRequest) *UpdateContactInvoker {
	requestDef := GenReqDefForUpdateContact()
	return &UpdateContactInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCorp SP管理员修改企业
//
// 修改企业，若任一参数为null或者不携带则不修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateCorp(request *model.UpdateCorpRequest) (*model.UpdateCorpResponse, error) {
	requestDef := GenReqDefForUpdateCorp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCorpResponse), nil
	}
}

// UpdateCorpInvoker SP管理员修改企业
func (c *MeetingClient) UpdateCorpInvoker(request *model.UpdateCorpRequest) *UpdateCorpInvoker {
	requestDef := GenReqDefForUpdateCorp()
	return &UpdateCorpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCorpBasicInfo 企业管理员修改企业注册信息
//
// 企业管理员通过该接口修改企业注册信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateCorpBasicInfo(request *model.UpdateCorpBasicInfoRequest) (*model.UpdateCorpBasicInfoResponse, error) {
	requestDef := GenReqDefForUpdateCorpBasicInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCorpBasicInfoResponse), nil
	}
}

// UpdateCorpBasicInfoInvoker 企业管理员修改企业注册信息
func (c *MeetingClient) UpdateCorpBasicInfoInvoker(request *model.UpdateCorpBasicInfoRequest) *UpdateCorpBasicInfoInvoker {
	requestDef := GenReqDefForUpdateCorpBasicInfo()
	return &UpdateCorpBasicInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDepartment 修改部门
//
// 企业管理员通过该接口修改部门。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateDepartment(request *model.UpdateDepartmentRequest) (*model.UpdateDepartmentResponse, error) {
	requestDef := GenReqDefForUpdateDepartment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDepartmentResponse), nil
	}
}

// UpdateDepartmentInvoker 修改部门
func (c *MeetingClient) UpdateDepartmentInvoker(request *model.UpdateDepartmentRequest) *UpdateDepartmentInvoker {
	requestDef := GenReqDefForUpdateDepartment()
	return &UpdateDepartmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDevice 修改终端
//
// 企业管理员通过该接口修改专业会议终端。
// &gt; 如果需要修改Ideahub、SmartRooms、智慧屏TV请使用[[修改用户](https://support.huaweicloud.com/api-meeting/meeting_21_0068.html)](tag:hws)[[修改用户](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0068.html)](tag:hk)接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateDevice(request *model.UpdateDeviceRequest) (*model.UpdateDeviceResponse, error) {
	requestDef := GenReqDefForUpdateDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceResponse), nil
	}
}

// UpdateDeviceInvoker 修改终端
func (c *MeetingClient) UpdateDeviceInvoker(request *model.UpdateDeviceRequest) *UpdateDeviceInvoker {
	requestDef := GenReqDefForUpdateDevice()
	return &UpdateDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMaterial 更新信息窗素材
//
// 更新信息窗素材。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateMaterial(request *model.UpdateMaterialRequest) (*model.UpdateMaterialResponse, error) {
	requestDef := GenReqDefForUpdateMaterial()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMaterialResponse), nil
	}
}

// UpdateMaterialInvoker 更新信息窗素材
func (c *MeetingClient) UpdateMaterialInvoker(request *model.UpdateMaterialRequest) *UpdateMaterialInvoker {
	requestDef := GenReqDefForUpdateMaterial()
	return &UpdateMaterialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMeeting 编辑预约会议
//
// 该接口用于修改已预约的会议。会议开始后，不能被修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateMeeting(request *model.UpdateMeetingRequest) (*model.UpdateMeetingResponse, error) {
	requestDef := GenReqDefForUpdateMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMeetingResponse), nil
	}
}

// UpdateMeetingInvoker 编辑预约会议
func (c *MeetingClient) UpdateMeetingInvoker(request *model.UpdateMeetingRequest) *UpdateMeetingInvoker {
	requestDef := GenReqDefForUpdateMeeting()
	return &UpdateMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMemberVmr 修改用会议室及个人会议ID信息
//
// 企业用户登录后可以修改分配给用户的云会议室及个人会议ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateMemberVmr(request *model.UpdateMemberVmrRequest) (*model.UpdateMemberVmrResponse, error) {
	requestDef := GenReqDefForUpdateMemberVmr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMemberVmrResponse), nil
	}
}

// UpdateMemberVmrInvoker 修改用会议室及个人会议ID信息
func (c *MeetingClient) UpdateMemberVmrInvoker(request *model.UpdateMemberVmrRequest) *UpdateMemberVmrInvoker {
	requestDef := GenReqDefForUpdateMemberVmr()
	return &UpdateMemberVmrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMyInfo 用户修改自己的信息
//
// 企业用户通过该接口修改自己的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateMyInfo(request *model.UpdateMyInfoRequest) (*model.UpdateMyInfoResponse, error) {
	requestDef := GenReqDefForUpdateMyInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMyInfoResponse), nil
	}
}

// UpdateMyInfoInvoker 用户修改自己的信息
func (c *MeetingClient) UpdateMyInfoInvoker(request *model.UpdateMyInfoRequest) *UpdateMyInfoInvoker {
	requestDef := GenReqDefForUpdateMyInfo()
	return &UpdateMyInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProgram 更新信息窗节目
//
// 更新信息窗节目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateProgram(request *model.UpdateProgramRequest) (*model.UpdateProgramResponse, error) {
	requestDef := GenReqDefForUpdateProgram()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProgramResponse), nil
	}
}

// UpdateProgramInvoker 更新信息窗节目
func (c *MeetingClient) UpdateProgramInvoker(request *model.UpdateProgramRequest) *UpdateProgramInvoker {
	requestDef := GenReqDefForUpdateProgram()
	return &UpdateProgramInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePublication 修改信息窗发布
//
// 修改信息窗发布。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdatePublication(request *model.UpdatePublicationRequest) (*model.UpdatePublicationResponse, error) {
	requestDef := GenReqDefForUpdatePublication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePublicationResponse), nil
	}
}

// UpdatePublicationInvoker 修改信息窗发布
func (c *MeetingClient) UpdatePublicationInvoker(request *model.UpdatePublicationRequest) *UpdatePublicationInvoker {
	requestDef := GenReqDefForUpdatePublication()
	return &UpdatePublicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePwd 修改密码
//
// 企业成员通过该接口提供用户修改密码功能，服务器收到请求，修改用户密码并返回结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdatePwd(request *model.UpdatePwdRequest) (*model.UpdatePwdResponse, error) {
	requestDef := GenReqDefForUpdatePwd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePwdResponse), nil
	}
}

// UpdatePwdInvoker 修改密码
func (c *MeetingClient) UpdatePwdInvoker(request *model.UpdatePwdRequest) *UpdatePwdInvoker {
	requestDef := GenReqDefForUpdatePwd()
	return &UpdatePwdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRecurringMeeting 编辑周期性会议
//
// 该接口用于修改已预约的周期性会议。会议开始后，不能被修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateRecurringMeeting(request *model.UpdateRecurringMeetingRequest) (*model.UpdateRecurringMeetingResponse, error) {
	requestDef := GenReqDefForUpdateRecurringMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecurringMeetingResponse), nil
	}
}

// UpdateRecurringMeetingInvoker 编辑周期性会议
func (c *MeetingClient) UpdateRecurringMeetingInvoker(request *model.UpdateRecurringMeetingRequest) *UpdateRecurringMeetingInvoker {
	requestDef := GenReqDefForUpdateRecurringMeeting()
	return &UpdateRecurringMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRecurringSubMeeting 编辑周期性会议的子会议
//
// 该接口用于修改已预约的周期性会议的子会议。会议开始后，不能被修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateRecurringSubMeeting(request *model.UpdateRecurringSubMeetingRequest) (*model.UpdateRecurringSubMeetingResponse, error) {
	requestDef := GenReqDefForUpdateRecurringSubMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecurringSubMeetingResponse), nil
	}
}

// UpdateRecurringSubMeetingInvoker 编辑周期性会议的子会议
func (c *MeetingClient) UpdateRecurringSubMeetingInvoker(request *model.UpdateRecurringSubMeetingRequest) *UpdateRecurringSubMeetingInvoker {
	requestDef := GenReqDefForUpdateRecurringSubMeeting()
	return &UpdateRecurringSubMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResource SP管理员根据修改企业资源
//
// 企业修改资源的过期时间、停用状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateResource(request *model.UpdateResourceRequest) (*model.UpdateResourceResponse, error) {
	requestDef := GenReqDefForUpdateResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResourceResponse), nil
	}
}

// UpdateResourceInvoker SP管理员根据修改企业资源
func (c *MeetingClient) UpdateResourceInvoker(request *model.UpdateResourceRequest) *UpdateResourceInvoker {
	requestDef := GenReqDefForUpdateResource()
	return &UpdateResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRoomSetting 修改网络研讨会高级设置
//
// 该接口用于设置指定网络研讨会的高级设置。管理员可设置企业内的网络研讨会高级设置，非管理员只可设置自己预定的网络研讨会的高级设置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateRoomSetting(request *model.UpdateRoomSettingRequest) (*model.UpdateRoomSettingResponse, error) {
	requestDef := GenReqDefForUpdateRoomSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRoomSettingResponse), nil
	}
}

// UpdateRoomSettingInvoker 修改网络研讨会高级设置
func (c *MeetingClient) UpdateRoomSettingInvoker(request *model.UpdateRoomSettingRequest) *UpdateRoomSettingInvoker {
	requestDef := GenReqDefForUpdateRoomSetting()
	return &UpdateRoomSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStartedConfConfig 会中修改配置
//
// 该接口用于修改会议配置，包括会议共享是否锁定，允许呼入范围。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateStartedConfConfig(request *model.UpdateStartedConfConfigRequest) (*model.UpdateStartedConfConfigResponse, error) {
	requestDef := GenReqDefForUpdateStartedConfConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStartedConfConfigResponse), nil
	}
}

// UpdateStartedConfConfigInvoker 会中修改配置
func (c *MeetingClient) UpdateStartedConfConfigInvoker(request *model.UpdateStartedConfConfigRequest) *UpdateStartedConfConfigInvoker {
	requestDef := GenReqDefForUpdateStartedConfConfig()
	return &UpdateStartedConfConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateToken 刷新Token
//
// 该接口提供刷新Token功能，根据传入的Token，刷新Token失效时间并返回结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateToken(request *model.UpdateTokenRequest) (*model.UpdateTokenResponse, error) {
	requestDef := GenReqDefForUpdateToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTokenResponse), nil
	}
}

// UpdateTokenInvoker 刷新Token
func (c *MeetingClient) UpdateTokenInvoker(request *model.UpdateTokenRequest) *UpdateTokenInvoker {
	requestDef := GenReqDefForUpdateToken()
	return &UpdateTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUser 修改用户
//
// 企业管理员通过该接口修改企业用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateUser(request *model.UpdateUserRequest) (*model.UpdateUserResponse, error) {
	requestDef := GenReqDefForUpdateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserResponse), nil
	}
}

// UpdateUserInvoker 修改用户
func (c *MeetingClient) UpdateUserInvoker(request *model.UpdateUserRequest) *UpdateUserInvoker {
	requestDef := GenReqDefForUpdateUser()
	return &UpdateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWebHookConfigStatus 开启事件推送
//
// 该接口用于管理员变更订阅配置使用状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateWebHookConfigStatus(request *model.UpdateWebHookConfigStatusRequest) (*model.UpdateWebHookConfigStatusResponse, error) {
	requestDef := GenReqDefForUpdateWebHookConfigStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWebHookConfigStatusResponse), nil
	}
}

// UpdateWebHookConfigStatusInvoker 开启事件推送
func (c *MeetingClient) UpdateWebHookConfigStatusInvoker(request *model.UpdateWebHookConfigStatusRequest) *UpdateWebHookConfigStatusInvoker {
	requestDef := GenReqDefForUpdateWebHookConfigStatus()
	return &UpdateWebHookConfigStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWebinar 编辑网络研讨会
//
// 该接口用于修改已创建的网络研讨会。网络研讨会开始后不能修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UpdateWebinar(request *model.UpdateWebinarRequest) (*model.UpdateWebinarResponse, error) {
	requestDef := GenReqDefForUpdateWebinar()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWebinarResponse), nil
	}
}

// UpdateWebinarInvoker 编辑网络研讨会
func (c *MeetingClient) UpdateWebinarInvoker(request *model.UpdateWebinarRequest) *UpdateWebinarInvoker {
	requestDef := GenReqDefForUpdateWebinar()
	return &UpdateWebinarInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadFile 上传图片
//
// 该接口用户上传网络研讨会高级设置用的图片。图片可用于网络研讨会的封面和Logo。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) UploadFile(request *model.UploadFileRequest) (*model.UploadFileResponse, error) {
	requestDef := GenReqDefForUploadFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadFileResponse), nil
	}
}

// UploadFileInvoker 上传图片
func (c *MeetingClient) UploadFileInvoker(request *model.UploadFileRequest) *UploadFileInvoker {
	requestDef := GenReqDefForUploadFile()
	return &UploadFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchQosHistoryMeetings 查询QoS历史会议列表
//
// 该接口用于查询企业内历史会议的QoS告警。仅旗舰版企业/标准版企业的企业管理员有权限查询。可以查询最近3个月内的数据。
// &gt; 仪表盘的QoS统计功能需要申请才能开通。请联系华为销售人员，并提供华为云会议企业ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchQosHistoryMeetings(request *model.SearchQosHistoryMeetingsRequest) (*model.SearchQosHistoryMeetingsResponse, error) {
	requestDef := GenReqDefForSearchQosHistoryMeetings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchQosHistoryMeetingsResponse), nil
	}
}

// SearchQosHistoryMeetingsInvoker 查询QoS历史会议列表
func (c *MeetingClient) SearchQosHistoryMeetingsInvoker(request *model.SearchQosHistoryMeetingsRequest) *SearchQosHistoryMeetingsInvoker {
	requestDef := GenReqDefForSearchQosHistoryMeetings()
	return &SearchQosHistoryMeetingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchQosOnlineMeetings 查询QoS在线会议列表
//
// 该接口用于查询企业内正在召开会议的QoS告警。仅旗舰版企业/标准版企业的企业管理员有权限查询。
// &gt; 仪表盘的QoS统计功能需要申请才能开通。请联系华为销售人员，并提供华为云会议企业ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchQosOnlineMeetings(request *model.SearchQosOnlineMeetingsRequest) (*model.SearchQosOnlineMeetingsResponse, error) {
	requestDef := GenReqDefForSearchQosOnlineMeetings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchQosOnlineMeetingsResponse), nil
	}
}

// SearchQosOnlineMeetingsInvoker 查询QoS在线会议列表
func (c *MeetingClient) SearchQosOnlineMeetingsInvoker(request *model.SearchQosOnlineMeetingsRequest) *SearchQosOnlineMeetingsInvoker {
	requestDef := GenReqDefForSearchQosOnlineMeetings()
	return &SearchQosOnlineMeetingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchQosParticipantDetail 查询与会者的QoS数据
//
// 该接口用于查询企业内在线会议或历史会议的与会者QoS数据。仅旗舰版企业/标准版企业的企业管理员有权限查询。
// &gt; 仪表盘的QoS统计功能需要申请才能开通。请联系华为销售人员，并提供华为云会议企业ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchQosParticipantDetail(request *model.SearchQosParticipantDetailRequest) (*model.SearchQosParticipantDetailResponse, error) {
	requestDef := GenReqDefForSearchQosParticipantDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchQosParticipantDetailResponse), nil
	}
}

// SearchQosParticipantDetailInvoker 查询与会者的QoS数据
func (c *MeetingClient) SearchQosParticipantDetailInvoker(request *model.SearchQosParticipantDetailRequest) *SearchQosParticipantDetailInvoker {
	requestDef := GenReqDefForSearchQosParticipantDetail()
	return &SearchQosParticipantDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchQosParticipants 查询QoS会议与会者列表
//
// 该接口用于查询企业内在线会议或历史会议的与会者QoS告警。仅旗舰版企业/标准版企业的企业管理员有权限查询。
// &gt; 仪表盘的QoS统计功能需要申请才能开通。请联系华为销售人员，并提供华为云会议企业ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchQosParticipants(request *model.SearchQosParticipantsRequest) (*model.SearchQosParticipantsResponse, error) {
	requestDef := GenReqDefForSearchQosParticipants()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchQosParticipantsResponse), nil
	}
}

// SearchQosParticipantsInvoker 查询QoS会议与会者列表
func (c *MeetingClient) SearchQosParticipantsInvoker(request *model.SearchQosParticipantsRequest) *SearchQosParticipantsInvoker {
	requestDef := GenReqDefForSearchQosParticipants()
	return &SearchQosParticipantsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetQosThreshold 设置QoS告警阈值
//
// 该接口用于设置QoS告警阈值。仅旗舰版企业/标准版企业的企业管理员有权限设置。
// &gt; 仪表盘的QoS统计功能需要申请才能开通。请联系华为销售人员，并提供华为云会议企业ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SetQosThreshold(request *model.SetQosThresholdRequest) (*model.SetQosThresholdResponse, error) {
	requestDef := GenReqDefForSetQosThreshold()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetQosThresholdResponse), nil
	}
}

// SetQosThresholdInvoker 设置QoS告警阈值
func (c *MeetingClient) SetQosThresholdInvoker(request *model.SetQosThresholdRequest) *SetQosThresholdInvoker {
	requestDef := GenReqDefForSetQosThreshold()
	return &SetQosThresholdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQosThreshold 查询QoS告警阈值
//
// 该接口用于查询QoS告警阈值。仅旗舰版企业/标准版企业的企业管理员有权限查询。
// &gt; 该接口用于查询QoS告警阈值。仅旗舰版企业/标准版企业的企业管理员有权限查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) ShowQosThreshold(request *model.ShowQosThresholdRequest) (*model.ShowQosThresholdResponse, error) {
	requestDef := GenReqDefForShowQosThreshold()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQosThresholdResponse), nil
	}
}

// ShowQosThresholdInvoker 查询QoS告警阈值
func (c *MeetingClient) ShowQosThresholdInvoker(request *model.ShowQosThresholdRequest) *ShowQosThresholdInvoker {
	requestDef := GenReqDefForShowQosThreshold()
	return &ShowQosThresholdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchStatisticConferenceInfo 查询企业级会议总体统计数据
//
// 该接口用于查询企业内：
// * 单日内按小时统计的会议数据。
// * 指定日期范围内按日/按月统计的会议数据。
// &gt; 仅旗舰版企业/标准版企业的企业管理员有权限查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchStatisticConferenceInfo(request *model.SearchStatisticConferenceInfoRequest) (*model.SearchStatisticConferenceInfoResponse, error) {
	requestDef := GenReqDefForSearchStatisticConferenceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchStatisticConferenceInfoResponse), nil
	}
}

// SearchStatisticConferenceInfoInvoker 查询企业级会议总体统计数据
func (c *MeetingClient) SearchStatisticConferenceInfoInvoker(request *model.SearchStatisticConferenceInfoRequest) *SearchStatisticConferenceInfoInvoker {
	requestDef := GenReqDefForSearchStatisticConferenceInfo()
	return &SearchStatisticConferenceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchStatisticConferenceParticipant 查询企业级会议与会统计数据
//
// 该接口用于查询企业内与会者数据统计：
// * 查询与会用户统计数据，按日/按月统计。
// * 查询与会硬件终端统计数据，按日/按月统计。
// * 查询与会设备统计数据，按日/按月统计。
// &gt; 仅旗舰版企业/标准版企业的企业管理员有权限查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchStatisticConferenceParticipant(request *model.SearchStatisticConferenceParticipantRequest) (*model.SearchStatisticConferenceParticipantResponse, error) {
	requestDef := GenReqDefForSearchStatisticConferenceParticipant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchStatisticConferenceParticipantResponse), nil
	}
}

// SearchStatisticConferenceParticipantInvoker 查询企业级会议与会统计数据
func (c *MeetingClient) SearchStatisticConferenceParticipantInvoker(request *model.SearchStatisticConferenceParticipantRequest) *SearchStatisticConferenceParticipantInvoker {
	requestDef := GenReqDefForSearchStatisticConferenceParticipant()
	return &SearchStatisticConferenceParticipantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchStatisticResourceInfo 查询企业级会议已购资源使用统计数据
//
// 该接口用于查询企业内已购资源使用状况数据统计：
// * 查询已购资源使用状况，按日/按月统计。
// &gt; 仅旗舰版企业/标准版企业的企业管理员有权限查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchStatisticResourceInfo(request *model.SearchStatisticResourceInfoRequest) (*model.SearchStatisticResourceInfoResponse, error) {
	requestDef := GenReqDefForSearchStatisticResourceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchStatisticResourceInfoResponse), nil
	}
}

// SearchStatisticResourceInfoInvoker 查询企业级会议已购资源使用统计数据
func (c *MeetingClient) SearchStatisticResourceInfoInvoker(request *model.SearchStatisticResourceInfoRequest) *SearchStatisticResourceInfoInvoker {
	requestDef := GenReqDefForSearchStatisticResourceInfo()
	return &SearchStatisticResourceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchStatisticUserInfo 查询企业级会议的用户统计数据
//
// 该接口用于查询企业内用户数据统计：
// * 查询会议用户登录数据，按日/按月统计。
// * 查询会议用户激活数据，按日/按月统计。
// * 查询会议用户登录设备数据，按日/按月统计。
// &gt; 仅旗舰版企业/标准版企业的企业管理员有权限查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MeetingClient) SearchStatisticUserInfo(request *model.SearchStatisticUserInfoRequest) (*model.SearchStatisticUserInfoResponse, error) {
	requestDef := GenReqDefForSearchStatisticUserInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchStatisticUserInfoResponse), nil
	}
}

// SearchStatisticUserInfoInvoker 查询企业级会议的用户统计数据
func (c *MeetingClient) SearchStatisticUserInfoInvoker(request *model.SearchStatisticUserInfoRequest) *SearchStatisticUserInfoInvoker {
	requestDef := GenReqDefForSearchStatisticUserInfo()
	return &SearchStatisticUserInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
