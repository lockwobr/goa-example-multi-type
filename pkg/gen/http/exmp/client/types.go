// Code generated by goa v3.5.2, DO NOT EDIT.
//
// exmp HTTP client types
//
// Command:
// $ goa gen github.com/lockwobr/goa-example-multi-type/pkg/design -o pkg

package client

import (
	exmpviews "github.com/lockwobr/goa-example-multi-type/pkg/gen/exmp/views"
)

// GetPersonResponseBody is the type of the "exmp" service "get-person"
// endpoint HTTP response body.
type GetPersonResponseBody struct {
	Name *string `json:"first_name"`
	Age  *int    `form:"age,omitempty" json:"age,omitempty" xml:"age,omitempty"`
}

// GetFamResponseBody is the type of the "exmp" service "get-fam" endpoint HTTP
// response body.
type GetFamResponseBody struct {
	Surname *string               `form:"surname,omitempty" json:"surname,omitempty" xml:"surname,omitempty"`
	Size    *int                  `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	Mother  *PersonResponseBody   `json:"mom"`
	Father  *PersonResponseBody   `json:"dad"`
	Kids    []*PersonResponseBody `form:"kids,omitempty" json:"kids,omitempty" xml:"kids,omitempty"`
}

// PersonResponseBody is used to define fields on response body types.
type PersonResponseBody struct {
	Name *string `json:"first_name"`
	Age  *int    `form:"age,omitempty" json:"age,omitempty" xml:"age,omitempty"`
}

// NewGetPersonPersonOK builds a "exmp" service "get-person" endpoint result
// from a HTTP "OK" response.
func NewGetPersonPersonOK(body *GetPersonResponseBody) *exmpviews.PersonView {
	v := &exmpviews.PersonView{
		Name: body.Name,
		Age:  body.Age,
	}

	return v
}

// NewGetFamFamOK builds a "exmp" service "get-fam" endpoint result from a HTTP
// "OK" response.
func NewGetFamFamOK(body *GetFamResponseBody) *exmpviews.FamView {
	v := &exmpviews.FamView{
		Surname: body.Surname,
		Size:    body.Size,
	}
	if body.Mother != nil {
		v.Mother = unmarshalPersonResponseBodyToExmpviewsPersonView(body.Mother)
	}
	if body.Father != nil {
		v.Father = unmarshalPersonResponseBodyToExmpviewsPersonView(body.Father)
	}
	if body.Kids != nil {
		v.Kids = make([]*exmpviews.PersonView, len(body.Kids))
		for i, val := range body.Kids {
			v.Kids[i] = unmarshalPersonResponseBodyToExmpviewsPersonView(val)
		}
	}

	return v
}