# ==================================================================================== #
# ENV
# ==================================================================================== #
GOSEC_VERSION?=v2.11.0
GOLINT_VERSION?=v1.44.2
GOBIN?=$(GOPATH)/bin
GOSEC?=$(GOBIN)/gosec

# ==================================================================================== #
# LINTERS
# ==================================================================================== #

.PHONY: lint/gosec
lint/gosec: ## Lint code with gosec
	@if [ -z $(shell command -v ${GOBIN}/gosec 2>/dev/null) ]; then ${MAKE} install/gosec; fi;
	$(call log_msg,"Running Golang security checker...")
	${GOBIN}/gosec --version
	${GOBIN}/gosec -fmt=$(fmt) -out=$(out) -exclude-dir=.go -log=/dev/null ./....PHONY: lint/gosec

.PHONY: lint/cli
lint/cli: ## Lint code with golang-cli lint
	@if [ -z $(shell command -v ${GOBIN}/golangci-lint 2>/dev/null) ]; then ${MAKE} install/lint; fi;
	$(call log_msg,"Running linters...")
	${GOBIN}/golangci-lint --version
	${GOBIN}/golangci-lint run

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

.PHONY: install/gosec
install/gosec: ## Install gosec
	@$(call log_msg,"Golang security checker installation...")
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b ${GOBIN} ${GOSEC_VERSION}

.PHONY: install/lint
install/lint: ## Install linter
	@$(call log_msg,"Linter installation...")
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOBIN} ${GOLINT_VERSION}

