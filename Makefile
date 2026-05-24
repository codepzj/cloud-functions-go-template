.PHONY: help run fmt

help: # Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

run: # Run the application
	set -a && . ./cloud-functions/.env && set +a && edgeone pages dev

fmt: # Format code
	cd cloud-functions && go fmt ./...