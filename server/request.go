package server

type InsertWordRequest struct {
	Words map[string]int32 `json:"words"`
}
