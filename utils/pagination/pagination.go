package pagination

type pagination struct {
}

type Pagination interface {
	GetPagination(offset, limit, totalData int) PaginationData
}

func NewPagination() Pagination {
	return &pagination{}
}

func (p *pagination) GetPagination(offset, limit, totalData int) PaginationData {
	currentPage := offset/limit + 1
	totalPage := totalData/limit + 1
	pageList := make([]int, totalPage)
	for i := 0; i < totalPage; i++ {
		pageList[i] = i + 1
	}

	return PaginationData{
		PreviousPage: currentPage - 1,
		CurrentPage:  currentPage,
		NextPage:     currentPage + 1,
		TotalPage:    totalPage,
		PageList:     pageList,
	}
}
