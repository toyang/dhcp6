// Code generated by "stringer -output=string.go -type=ArchType,DUIDType,MessageType,Status,OptionCode"; DO NOT EDIT.

package dhcp6

import "strconv"

const _ArchType_name = "ArchTypeIntelx86PCArchTypeNECPC98ArchTypeEFIItaniumArchTypeDECAlphaArchtypeArcx86ArchTypeIntelLeanClientArchTypeEFIIA32ArchTypeEFIBCArchTypeEFIXscaleArchTypeEFIx8664"

var _ArchType_index = [...]uint8{0, 18, 33, 51, 67, 81, 104, 119, 132, 149, 165}

func (i ArchType) String() string {
	if i >= ArchType(len(_ArchType_index)-1) {
		return "ArchType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ArchType_name[_ArchType_index[i]:_ArchType_index[i+1]]
}

const _DUIDType_name = "DUIDTypeLLTDUIDTypeENDUIDTypeLLDUIDTypeUUID"

var _DUIDType_index = [...]uint8{0, 11, 21, 31, 43}

func (i DUIDType) String() string {
	i -= 1
	if i >= DUIDType(len(_DUIDType_index)-1) {
		return "DUIDType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _DUIDType_name[_DUIDType_index[i]:_DUIDType_index[i+1]]
}

const _MessageType_name = "MessageTypeSolicitMessageTypeAdvertiseMessageTypeRequestMessageTypeConfirmMessageTypeRenewMessageTypeRebindMessageTypeReplyMessageTypeReleaseMessageTypeDeclineMessageTypeReconfigureMessageTypeInformationRequestMessageTypeRelayForwMessageTypeRelayReplMessageTypeLeasequeryMessageTypeLeasequeryReplyMessageTypeLeasequeryDoneMessageTypeLeasequeryDataMessageTypeReconfigureRequestMessageTypeReconfigureReplyMessageTypeDHCPv4QueryMessageTypeDHCPv4Response"

var _MessageType_index = [...]uint16{0, 18, 38, 56, 74, 90, 107, 123, 141, 159, 181, 210, 230, 250, 271, 297, 322, 347, 376, 403, 425, 450}

func (i MessageType) String() string {
	i -= 1
	if i >= MessageType(len(_MessageType_index)-1) {
		return "MessageType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _MessageType_name[_MessageType_index[i]:_MessageType_index[i+1]]
}

const _Status_name = "StatusSuccessStatusUnspecFailStatusNoAddrsAvailStatusNoBindingStatusNotOnLinkStatusUseMulticastStatusNoPrefixAvailStatusUnknownQueryTypeStatusMalformedQueryStatusNotConfiguredStatusNotAllowedStatusQueryTerminated"

var _Status_index = [...]uint8{0, 13, 29, 47, 62, 77, 95, 114, 136, 156, 175, 191, 212}

func (i Status) String() string {
	if i >= Status(len(_Status_index)-1) {
		return "Status(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}

const (
	_OptionCode_name_0 = "OptionClientIDOptionServerIDOptionIANAOptionIATAOptionIAAddrOptionOROOptionPreferenceOptionElapsedTimeOptionRelayMsg"
	_OptionCode_name_1 = "OptionAuthOptionUnicastOptionStatusCodeOptionRapidCommitOptionUserClassOptionVendorClassOptionVendorOptsOptionInterfaceIDOptionReconfMsgOptionReconfAccept"
	_OptionCode_name_2 = "OptionIAPDOptionIAPrefix"
	_OptionCode_name_3 = "OptionRemoteIdentifier"
	_OptionCode_name_4 = "OptionBootFileURLOptionBootFileParamOptionClientArchTypeOptionNII"
)

var (
	_OptionCode_index_0 = [...]uint8{0, 14, 28, 38, 48, 60, 69, 85, 102, 116}
	_OptionCode_index_1 = [...]uint8{0, 10, 23, 39, 56, 71, 88, 104, 121, 136, 154}
	_OptionCode_index_2 = [...]uint8{0, 10, 24}
	_OptionCode_index_3 = [...]uint8{0, 22}
	_OptionCode_index_4 = [...]uint8{0, 17, 36, 56, 65}
)

func (i OptionCode) String() string {
	switch {
	case 1 <= i && i <= 9:
		i -= 1
		return _OptionCode_name_0[_OptionCode_index_0[i]:_OptionCode_index_0[i+1]]
	case 11 <= i && i <= 20:
		i -= 11
		return _OptionCode_name_1[_OptionCode_index_1[i]:_OptionCode_index_1[i+1]]
	case 25 <= i && i <= 26:
		i -= 25
		return _OptionCode_name_2[_OptionCode_index_2[i]:_OptionCode_index_2[i+1]]
	case i == 37:
		return _OptionCode_name_3
	case 59 <= i && i <= 62:
		i -= 59
		return _OptionCode_name_4[_OptionCode_index_4[i]:_OptionCode_index_4[i+1]]
	default:
		return "OptionCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
