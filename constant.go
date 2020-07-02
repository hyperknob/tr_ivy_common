package tr_ivy_common

/**
 * 错误码返回格式
 */
type ErrResult struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Err string `json:"err"`
}

/**
 * 错误码定义
 */

var (
	ERR_SIGN_NO_MATCH            = ErrResult{10002, "failed", "sign no match!"}
	ERR_VERIFY_CODE_NO_MATCH     = ErrResult{10003, "failed", "verify code error!"}
	ERR_JWT_NO_MATCH             = ErrResult{10004, "failed", "jwt not correct!"}
	ERR_MOBILE_ALREADY_EXISTED   = ErrResult{10005, "failed", "mobile number already bound, pls unbound firstly!"}
	ERR_SEND_VERIFY_CODE_ERR     = ErrResult{10006, "failed", "send verification code error!"}
	ERR_VALIDATE_VERIFY_CODE_ERR = ErrResult{10007, "failed", "validate verification code error!"}
	ERR_NEED_BIND_MOBILE         = ErrResult{10008, "failed", "need bind mobile phone!"}
	ERR_PHOTO_EXT_NAME_WRONG     = ErrResult{10009, "failed", "photo ext file name err!"}
	ERR_PHOTO_SIZE_TOO_BIG       = ErrResult{10010, "failed", "photo size bigger than 5MB!"}
	ERR_PHOTO_UPLOAD_FAILED      = ErrResult{10011, "failed", "photo upload failed!"}
	ERR_WIDGET_IS_NULL           = ErrResult{10012, "failed", "avatar widget is null!"}
	ERR_DEVICE_ID_IS_NULL        = ErrResult{10013, "failed", "deviceId is null!"}
	ERR_DEVICE_PLATFORM_IS_NULL  = ErrResult{10014, "failed", "X-Platform is null!"}

	ERR_OWNER_CATEGORY_IS_NULL  = ErrResult{20001, "failed", "owner category is null!"}
	ERR_FIND_CATEGORY_IS_NULL  = ErrResult{20002, "failed", "find category is null!"}
	ERR_IVY_SUBID_IS_NULL  = ErrResult{20003, "failed", "ivySubId is null!"}
	ERR_ALBUM_ID_IS_NULL  = ErrResult{20004, "failed", "albumId is null!"}
	ERR_IVY_OWNER_UID_IS_NULL  = ErrResult{20004, "failed", "ivyOwnerUid is null!"}
)
