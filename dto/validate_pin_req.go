package dto

type ValidatePinReq struct {
	PIN    string `json:"pin"`
	UserId string `json: "user_id"`
}
