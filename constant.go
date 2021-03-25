package tr_ivy_common

/**
 * 错误码返回格式
 */
type ErrResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Err  string `json:"err"`
	CnErr string `json:"cnErr"`
}

/**
 * 错误码定义
 */

var (
	ERR_SIGN_NO_MATCH                       = ErrResult{10002, "failed", "sign no match!","API签名验证失败"}
	ERR_VERIFY_CODE_NO_MATCH                = ErrResult{10003, "failed", "verify code error!","验证码错误"}
	ERR_JWT_NO_MATCH                        = ErrResult{10004, "failed", "jwt not correct!", "用户令牌错误"}
	ERR_MOBILE_ALREADY_EXISTED              = ErrResult{10005, "failed", "mobile number already bound, pls unbound firstly!", "手机号已绑定，请首先解绑"}
	ERR_SEND_VERIFY_CODE_ERR                = ErrResult{10006, "failed", "send verification code error!", "发送验证码失败"}
	ERR_VALIDATE_VERIFY_CODE_ERR            = ErrResult{10007, "failed", "validate verification code error!", "验证码验证错误"}
	ERR_NEED_BIND_MOBILE                    = ErrResult{10008, "failed", "need bind mobile phone!","请注册/登录"}
	ERR_PHOTO_EXT_NAME_WRONG                = ErrResult{10009, "failed", "photo ext file name err!","图片扩展名不匹配"}
	ERR_PHOTO_SIZE_TOO_BIG                  = ErrResult{10010, "failed", "photo size bigger than 5MB!","图片尺寸不能超过5MB"}
	ERR_PHOTO_UPLOAD_FAILED                 = ErrResult{10011, "failed", "photo upload failed!","图片上传错误"}
	ERR_WIDGET_IS_NULL                      = ErrResult{10012, "failed", "avatar widget is null!","头像挂件为空"}
	ERR_DEVICE_ID_IS_NULL                   = ErrResult{10013, "failed", "deviceId is null!","设备标识为空"}
	ERR_DEVICE_PLATFORM_IS_NULL             = ErrResult{10014, "failed", "X-Platform is null!","平台标识为空"}
	ERR_PHOTO_DECODE_FAIL                   = ErrResult{10015, "failed", "photo decode failed!","图片解码失败"}
	ERR_MKDIR_FAIL                          = ErrResult{10016, "failed", "make dir failed!","创建目录失败"}
	ERR_WRITE_FILE_FAIL                     = ErrResult{10017, "failed", "write file failed!","写入文件失败"}
	ERR_NEED_BIND_MOBILE_AND_IDENTIFICATION = ErrResult{10018, "failed", "need bind mobile phone and identification!","请注册/登录并实名验证"}
	ERR_NEED_IDENTIFICATION                 = ErrResult{10019, "failed", "need identification!","请实名验证"}
	ERR_IDENTIFICATION_FAIL                 = ErrResult{10020, "failed", "identification! failed","实名验证错误"}

	ERR_OWNER_CATEGORY_IS_NULL        = ErrResult{20001, "failed", "owner category is null!","UP主分类为空"}
	ERR_FIND_CATEGORY_IS_NULL         = ErrResult{20002, "failed", "find category is null!","发现分类为空"}
	ERR_IVY_SUBID_IS_NULL             = ErrResult{20003, "failed", "ivySubId is null!","视簇标识为空"}
	ERR_ALBUM_ID_IS_NULL              = ErrResult{20004, "failed", "albumId is null!","合集标识为空"}
	ERR_IVY_OWNER_UID_IS_NULL         = ErrResult{20005, "failed", "ivyOwnerUid is null!","UP主的标识为空"}
	ERR_SEARCH_WORD_IS_NULL           = ErrResult{20006, "failed", "searchWord is null!", "搜索关键词为空"}
	ERR_VOICE_HOUSE_ID_IS_NULL        = ErrResult{20007, "failed", "voiceHouseId is null!","房间标识为空"}
	ERR_VOICE_HOUSE_OWNER_UID_IS_NULL = ErrResult{20008, "failed", "voiceHouseOwnerUid is null!", "房间主持人标识为空"}

	ERR_AGORA_GEN_RTC_TOKEN_FAILED = ErrResult{30001, "failed", "gen rtc token failed!","生成RTC令牌失败"}
	ERR_AGORA_GEN_RTM_TOKEN_FAILED = ErrResult{30002, "failed", "gen rtm token failed!","生成RTM令牌失败"}
)

// 用于做major_color_tone颜色转换的map（由于python算法计算的主色调不合格，因此需要微调做个映射）
var MAJOR_COLOR_MAP = map[string]string{
	"":        "#F5D57F",
	"#000000": "#C1E0BC",
	"#666666": "#7DA9B8",
	"#808A87": "#8BCCB8",
	"#808069": "#B8B88C",
	"#E6E6E6": "#E67878",
	"#F5F5F5": "#F5BA93",
	"#FCE6C9": "#FCD097",
	"#FF0000": "#F54545",
	"#E3170D": "#E0463F",
	"#9C661F": "#E09B3F",
	"#FF7F50": "#FF7F50",
	"#FF6347": "#FF6347",
	"#FFC0CB": "#FF859A",
	"#B0171F": "#B0171F",
	"#FF00FF": "#E075E0",
	"#990033": "#990033",
	"#00FF00": "#E075E0",
	"#00FFFF": "#58F5F5",
	"#7FFF00": "#A1F54E",
	"#40E0D0": "#40E0D0",
	"#082E54": "#145BA3",
	"#228B22": "#55A355",
	"#6B8E23": "#6B8E23",
	"#0000FF": "#6C6CF5",
	"#03A89E": "#07B8AD",
	"#191970": "#1D1D8F",
	"#00C78C": "#38C79D",
	"#FFFFFF": "#A6E2F5",
	"#F0FFFF": "#58F5F5",
	"#CCCCCC": "#CC6A6A",
	"#FAFFF0": "#D4F593",
	"#FAF0E6": "#F5B06C",
	"#FFFFCD": "#F5F57F",
	"#FFF5EE": "#F5B080",
	"#FFFF00": "#F5F558",
	"#FF9912": "#FF9912",
	"#E3CF57": "#E3CF57",
	"#FFD700": "#FFE247",
	"#FF7D40": "#FF7D40",
	"#FFE384": "#FFE384",
	"#FF8000": "#FF8000",
	"#ED9121": "#ED9121",
	"#8B864E": "#CCC35A",
	"#802A2A": "#802A2A",
	"#C76114": "#CC4343",
	"#F4A460": "#F4A460",
	"#D2B48C": "#D2A05C",
	"#BC8F8F": "#CC5A5A",
	"#A0522D": "#B85E34",
	"#A020F0": "#BE87E0",
	"#DA70D6": "#DA83D7",
	"#8A2BE2": "#9B51E0",
	"#9933FA": "#6A2EA3",
}
