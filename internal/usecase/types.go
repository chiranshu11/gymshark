package usecase

type Result struct {
	Requested int         `json:"requested"`
	UsedSum   int         `json:"used_sum"`
	Overhead  int         `json:"overhead"`
	Packs     map[int]int `json:"packs"`
}
