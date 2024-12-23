package comments

import (
	. "stugi/api-gateway/internal/model"
)

type Service interface {
	AddComment(comment *Comment) error
}
