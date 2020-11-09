package response

type responseJson struct {
	Code    int         `json:"c"`
	Message string      `json:"m,omitempty"`
	Data    interface{} `json:"d,omitempty"`
}
