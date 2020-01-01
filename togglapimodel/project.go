package togglapimodel

type Project struct {
	Data data `json:"data"`
}

type data struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
