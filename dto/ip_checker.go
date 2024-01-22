package dto

type IpChecker struct {
	Status   string `json: "status"`
	Query    string `json: "query"`
	TimeZone string `json: "timezone"`
	Lat      int64  `json: "lat"`
	Lon      int64  `json: "lon"`
}
