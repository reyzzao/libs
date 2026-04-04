package valid_text_min

import (
	"testing"

	"github.com/reizzao/libs/feedback"
)

func Test_MinText(t *testing.T) {
	numTree := 3
	numFor := 4
	textOk := "foo"
	sutOK, _ := MinText(textOk, numTree)
	textFail := "f"
	_, sutFail := MinText(textFail, numFor)
	feedbackErrorMin := feedback.MinText(numFor)

	if sutOK != textOk {
		t.Error("Ops...Falhou o test MinText para Min")
	}

	if sutFail.Error() != feedbackErrorMin {
		t.Error("Ops...Falhou o test MinText para Min - Não deu o feedback de Error")
	}
}
