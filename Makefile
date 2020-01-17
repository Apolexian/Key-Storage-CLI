build:  # Build .go files into bin dir
	@echo "  > $(PROJECTNAME) is building go files."
	go build -o cmd/cli_vault cmd/main.go