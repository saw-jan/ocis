# standard colors
BLACK        := $(shell tput -Txterm setaf 0)
RED          := $(shell tput -Txterm setaf 1)
GREEN        := $(shell tput -Txterm setaf 2)
YELLOW       := $(shell tput -Txterm setaf 3)
LIGHTPURPLE  := $(shell tput -Txterm setaf 4)
PURPLE       := $(shell tput -Txterm setaf 5)
BLUE         := $(shell tput -Txterm setaf 6)
WHITE        := $(shell tput -Txterm setaf 7)

CUR_DIR = $(shell pwd)
OCIS_DIR = $(shell cd $(DIR)../../ && pwd)
BEHAT = $(OCIS_DIR)/vendor-bin/behat/vendor/bin/behat
BEHAT_YML = $(CUR_DIR)/config/behat.yml
BEHAT_SUITE ?= ""

.PHONY: test-paralleldeployment-api

test-paralleldeployment-api:
	@echo $(BLUE)"Running $(YELLOW)${BEHAT_SUITE}$(BLUE) tests tagged $(YELLOW)${BEHAT_FILTER_TAGS}$(BLUE):"$(WHITE)
	$(BEHAT) --colors -f pretty -f junit -c $(BEHAT_YML) --suite ${BEHAT_SUITE} --tags ${BEHAT_FILTER_TAGS}