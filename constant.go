package tr_ivy_common

/**
 * 错误码返回格式
 */
type ErrResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Err  string `json:"err"`
}

/**
 * 错误码定义
 */

var (
	ERR_SIGN_NO_MATCH                       = ErrResult{10002, "failed", "sign no match!"}
	ERR_VERIFY_CODE_NO_MATCH                = ErrResult{10003, "failed", "verify code error!"}
	ERR_JWT_NO_MATCH                        = ErrResult{10004, "failed", "jwt not correct!"}
	ERR_MOBILE_ALREADY_EXISTED              = ErrResult{10005, "failed", "mobile number already bound, pls unbound firstly!"}
	ERR_SEND_VERIFY_CODE_ERR                = ErrResult{10006, "failed", "send verification code error!"}
	ERR_VALIDATE_VERIFY_CODE_ERR            = ErrResult{10007, "failed", "validate verification code error!"}
	ERR_NEED_BIND_MOBILE                    = ErrResult{10008, "failed", "need bind mobile phone!"}
	ERR_PHOTO_EXT_NAME_WRONG                = ErrResult{10009, "failed", "photo ext file name err!"}
	ERR_PHOTO_SIZE_TOO_BIG                  = ErrResult{10010, "failed", "photo size bigger than 5MB!"}
	ERR_PHOTO_UPLOAD_FAILED                 = ErrResult{10011, "failed", "photo upload failed!"}
	ERR_WIDGET_IS_NULL                      = ErrResult{10012, "failed", "avatar widget is null!"}
	ERR_DEVICE_ID_IS_NULL                   = ErrResult{10013, "failed", "deviceId is null!"}
	ERR_DEVICE_PLATFORM_IS_NULL             = ErrResult{10014, "failed", "X-Platform is null!"}
	ERR_PHOTO_DECODE_FAIL                   = ErrResult{10015, "failed", "photo decode failed!"}
	ERR_MKDIR_FAIL                          = ErrResult{10016, "failed", "make dir failed!"}
	ERR_WRITE_FILE_FAIL                     = ErrResult{10017, "failed", "write file failed!"}
	ERR_NEED_BIND_MOBILE_AND_IDENTIFICATION = ErrResult{10018, "failed", "need bind mobile phone and identification!"}
	ERR_NEED_IDENTIFICATION                 = ErrResult{10019, "failed", "need identification!"}
	ERR_IDENTIFICATION_FAIL                 = ErrResult{10020, "failed", "identification! failed"}

	ERR_OWNER_CATEGORY_IS_NULL = ErrResult{20001, "failed", "owner category is null!"}
	ERR_FIND_CATEGORY_IS_NULL  = ErrResult{20002, "failed", "find category is null!"}
	ERR_IVY_SUBID_IS_NULL      = ErrResult{20003, "failed", "ivySubId is null!"}
	ERR_ALBUM_ID_IS_NULL       = ErrResult{20004, "failed", "albumId is null!"}
	ERR_IVY_OWNER_UID_IS_NULL  = ErrResult{20005, "failed", "ivyOwnerUid is null!"}
	ERR_SEARCH_WORD_IS_NULL    = ErrResult{20006, "failed", "searchWord is null!"}
	ERR_VOICE_HOUSE_ID_IS_NULL = ErrResult{20007, "failed", "voiceHouseId is null!"}

	ERR_AGORA_GEN_RTC_TOKEN_FAILED = ErrResult{30001, "failed", "gen rtc token failed!"}
	ERR_AGORA_GEN_RTM_TOKEN_FAILED = ErrResult{30002, "failed", "gen rtm token failed!"}
)

// 用于做major_color_tone颜色转换的map（由于python算法计算的主色调不合格，因此需要微调做个映射）
var MAJOR_COLOR_MAP = map[string]string {
	"" : "#F5D57F",
	"#000000" : "#C1E0BC",
	"#666666" : "#7DA9B8",
	"#808A87" : "#8BCCB8",
	"#808069" : "#B8B88C",
	"#E6E6E6" : "#E67878",
	"#F5F5F5" : "#F5BA93",
	"#FCE6C9" : "#FCD097",
	"#FF0000" : "#F54545",
	"#E3170D" : "#E0463F",
	"#9C661F" : "#E09B3F",
	"#FF7F50" : "#FF7F50",
	"#FF6347" : "#FF6347",
	"#FFC0CB" : "#FF859A",
	"#B0171F" : "#B0171F",
	"#FF00FF" : "#E075E0",
	"#990033" : "#990033",
	"#00FF00" : "#E075E0",
	"#00FFFF" : "#58F5F5",
	"#7FFF00" : "#A1F54E",
	"#40E0D0" : "#40E0D0",
	"#082E54" : "#145BA3",
	"#228B22" : "#55A355",
	"#6B8E23" : "#6B8E23",
	"#0000FF" : "#6C6CF5",
	"#03A89E" : "#07B8AD",
	"#191970" : "#1D1D8F",
	"#00C78C" : "#38C79D",
	"#FFFFFF" : "#A6E2F5",
	"#F0FFFF" : "#58F5F5",
	"#CCCCCC" : "#CC6A6A",
	"#FAFFF0" : "#D4F593",
	"#FAF0E6" : "#F5B06C",
	"#FFFFCD" : "#F5F57F",
	"#FFF5EE" : "#F5B080",
	"#FFFF00" : "#F5F558",
	"#FF9912" : "#FF9912",
	"#E3CF57" : "#E3CF57",
	"#FFD700" : "#FFE247",
	"#FF7D40" : "#FF7D40",
	"#FFE384" : "#FFE384",
	"#FF8000" : "#FF8000",
	"#ED9121" : "#ED9121",
	"#8B864E" : "#CCC35A",
	"#802A2A" : "#802A2A",
	"#C76114" : "#CC4343",
	"#F4A460" : "#F4A460",
	"#D2B48C" : "#D2A05C",
	"#BC8F8F" : "#CC5A5A",
	"#A0522D" : "#B85E34",
	"#A020F0" : "#BE87E0",
	"#DA70D6" : "#DA83D7",
	"#8A2BE2" : "#9B51E0",
	"#9933FA" : "#6A2EA3",
}
