package headers

const (
	// User agent client hints
	// ref. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#user_agent_client_hints
	SecChUa                = "Sec-Ch-Ua"
	SecChUaArch            = "Sec-Ch-Ua-Arch"
	SecChUaBitness         = "Sec-Ch-Ua-Bitness"
	SecChUaFullVersion     = "Sec-Ch-Ua-Full-Version"
	SecChUaFullVersionList = "Sec-Ch-Ua-Full-Version-List"
	SecChUaMobile          = "Sec-Ch-Ua-Mobile"
	SecChUaModel           = "Sec-Ch-Ua-Model"
	SecChUaPlatform        = "Sec-Ch-Ua-Platform"
	SecChUaPlatformVarsion = "Sec-Ch-Ua-Platform-Version"

	// (Deprecated) Device client hints
	// ref. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#device_client_hints
	ContentDPR    = "Content-Dpr"
	DeviceMemory  = "Device-Memory"
	DPR           = "Dpr"
	ViewportWidth = "Viewport-Width"
	Width         = "Width"

	// Network client hints
	// ref. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#network_client_hints
	Downlink = "Downlink"
	ECT      = "Ect"
	RTT      = "Rtt"
	SaveData = "Save-Data"
)
