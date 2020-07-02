package tr_ivy_common

// 获取复合设备标识
func GetDeviceComboId(deviceId string, devicePlatform string) (deviceComboId string) {
	if deviceId != "" && devicePlatform != "" {
		deviceComboId = deviceId + devicePlatform
		return
	} else {
		return ""
	}
}