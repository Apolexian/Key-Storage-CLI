build:  # Build .go files into bin dir
	@echo "  > $(PROJECTNAME) is building go files."
	go build -o cmd/api_storage cmd/main.go