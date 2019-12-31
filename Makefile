PROJECTNAME=$(shell basename "$(PWD)")
TEST_DIR = ../$(PROJECTNAME)/test/

.PHONY: test
test:   # Run tests from test dir
		@echo "  > $(PROJECTNAME) is running test files."
		go test $(TEST_DIR)

build:  # Build .go files into bin dir
	@echo "  > $(PROJECTNAME) is building go files."
	+$(MAKE) -C cmd
	+$(MAKE) -C internal

help:   # Show the help menu
	@fgrep -h "#" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
