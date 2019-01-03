package controller

// ClientCreateContainerRequest contains user request
type ClientCreateContainerRequest struct {
	BaseImage string `json:"baseImage"`
}

// ClientCreateContainerResponse stores response message for user
type ClientCreateContainerResponse struct {
	OK  bool   `json:"ok"`
	Msg string `json:"msg"`
}
