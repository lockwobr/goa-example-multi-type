package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("exmp", func() {
	Title("example Service")

	Server("exmp", func() {
		Services("exmp")
		Host("development", func() {
			URI("http://localhost:9000/")
		})

	})
})

var _ = Service("exmp", func() {

	Method("get-person", func() {
		HTTP(func() {
			GET("/person")
			Response(StatusOK)
		})
		Result(person)
	})
	Method("get-fam", func() {
		HTTP(func() {
			GET("/family")
			Response(StatusOK)
		})
		Result(family)
	})
})

var person = ResultType("application/person", "person", func() {
	Attribute("name", String, func() {
		Meta("struct:tag:json", "first_name")
	})
	Attribute("age", Int)
})

var family = ResultType("application/family", "fam", func() {
	Attribute("surname", String)
	Attribute("size", Int)
	Attribute("mother", person, func() {
		Meta("struct:tag:json", "mom") // if you comment one of these out it works, but having tags on both breaks the api output
	})
	Attribute("father", person, func() {
		Meta("struct:tag:json", "dad")
	})
	Attribute("kids", ArrayOf(person))
})
