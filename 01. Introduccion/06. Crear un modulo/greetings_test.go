package greetings

import "testing"
func TestHelloName(t *testing.T){
	name := "Juan"
	want := regexp.MustgCompile(`\b`+name+`\b`)

	msg, err := Hello("Juan")

	if !want.MatchString(msg) || err !=nil {
		t.Failf(`Hello("Juan) = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}

//go test -v