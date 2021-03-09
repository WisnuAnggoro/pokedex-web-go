package pagination

type PaginationData struct {
	PreviousPage int   `json:"previous_page"`
	CurrentPage  int   `json:"current_page"`
	NextPage     int   `json:"next_page"`
	TotalPage    int   `json:"total_page"`
	PageList     []int `json:"page_list"`
}
