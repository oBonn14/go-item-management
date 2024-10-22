package pagination

import (
	"gorm.io/gorm"
	"strings"
)

type Request struct {
	Page int
	Size int
}

type Pagination struct {
	TotalRecords int `json:"totalRecords"`
	Page         int `json:"page"`
	Size         int `json:"size"`
	TotalPages   int `json:"totalPages"`
}

func New(total int64, page, size int) *Pagination {
	totalRecords := int(total)

	totalPages := totalRecords / size
	if totalRecords%size != 0 {
		totalPages++
	}

	return &Pagination{
		TotalRecords: totalRecords,
		Page:         page,
		Size:         size,
		TotalPages:   totalPages,
	}
}

func Paginate(r *Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := r.Page
		if page <= 0 {
			page = 1
		}

		size := r.Size
		switch {
		case size > 100:
			size = 100
		case size < 0:
			size = 10
		}

		return db.Offset((page - 1) * size).Limit(size)
	}
}

func isStringColum(column string) bool {
	validStringColumn := map[string]bool{"" +
		"name": true,
		"age": true,
	}

	return validStringColumn[strings.ToLower(column)]
}
