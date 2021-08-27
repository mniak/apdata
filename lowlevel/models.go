package lowlevel

type GenericResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	SessionID string `json:"SessionID"`
	Message   string `json:"msg"`
	Hwd       int64  `json:"hwd"`
}
