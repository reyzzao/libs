package valid_text_max

import (
	"errors"

	"github.com/reizzao/libs/feedback"
)

func MaxText(text string, length int) (string, error) {
	if len(text) > length {
		return "", errors.New(feedback.MaxText(length))
	}
	return text, nil
}
