package calc

import (
	"context"
	"log"

	exmp "github.com/lockwobr/goa-example-multi-type/pkg/gen/exmp"
)

// exmp service example implementation.
// The example methods log the requests and return zero values.
type exmpsrvc struct {
	logger *log.Logger
}

// NewExmp returns the exmp service implementation.
func NewExmp(logger *log.Logger) exmp.Service {
	return &exmpsrvc{logger}
}

// GetPerson implements get-person.
func (s *exmpsrvc) GetPerson(ctx context.Context) (res *exmp.Person, err error) {
	s.logger.Print("exmp.get-person")

	res = person("Capiton Ron", 40)

	return
}

// GetFam implements get-fam.
func (s *exmpsrvc) GetFam(ctx context.Context) (res *exmp.Fam, err error) {
	s.logger.Print("exmp.get-fam")

	res = &exmp.Fam{
		Surname: sp("Harvey"),
		Size:    ip(4),
		Mother:  person("Katherine", 45),
		Father:  person("Martin", 50),
		Kids:    []*exmp.Person{person("Ben", 11), person("Caroline", 16)},
	}

	return
}

func person(name string, age int) *exmp.Person {
	return &exmp.Person{
		Name: &name,
		Age:  &age,
	}
}
func sp(s string) *string {
	return &s
}
func ip(s int) *int {
	return &s
}
