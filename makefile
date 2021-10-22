
name := github.com/lockwobr/goa-example-multi-type

build:
	go install $(name)/pkg/...

gen:
	goa gen $(name)/pkg/design -o pkg

exmp:
	goa example $(name)/pkg/design -o pkg