PROJECTNAME=$(shell basename "$(PWD)")

build:
	@echo "  > $(PROJECTNAME) is building."
	+$(MAKE) -C cmd
	+$(MAKE) -C internal
