package dto

type Station struct {
	Name  string `json:"name"`
	Forks [][]Station `json:"forks"`
	TrainColor string `json:"train_color"`
}
