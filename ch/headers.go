package ch

const (
	// User agent client hints
	// ref. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#user_agent_client_hints
	SecChUa                = "Sec-CH-UA"
	SecChUaArch            = "Sec-CH-UA-Arch"
	SecChUaBitness         = "Sec-CH-UA-Bitness"
	SecChUaFullVersion     = "Sec-CH-UA-Full-Version"
	SecChUaFullVersionList = "Sec-CH-UA-Full-Version-List"
	SecChUaMobile          = "Sec-CH-UA-Mobile"
	SecChUaModel           = "Sec-CH-UA-Model"
	SecChUaPlatform        = "Sec-CH-UA-Platform"
	SecChUaPlatformVarsion = "Sec-CH-UA-Platform-Version"

	// Device client hints
	// ref. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#device_client_hints
	ContentDPR    = "Content-DPR"
	DeviceMemory  = "Device-Memory"
	DPR           = "DPR"
	ViewportWidth = "Viewport-Width"
	Width         = "Width"

	// Network client hints
	// ref. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#network_client_hints
	Downlink = "Downlink"
	ECT      = "ECT"
	RTT      = "RTT"
	SaveData = "Save-Data"
)
