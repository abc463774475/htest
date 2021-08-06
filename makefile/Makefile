GOVER_MAJOR := $(shell go version | sed -E -e "s/.*go([0-9]+)[.]([0-9]+).*/\1/")
GOVER_MINOR := $(shell go version | sed -E -e "s/.*go([0-9]+)[.]([0-9]+).*/\2/")
g11 :=$(shell [ $(GOVER_MAJOR) -gt 1 ] || [ $(GOVER_MAJOR) -eq 1 ] && [ $(GOVER_MINOR) -ge 11 ]; echo $$?)

PACKAGE_LIST := echo $$(go list ./... | grep -vE "chaos-mesh/test|pkg/ptrace|zz_generated|vendor") github.com/chaos-mesh/chaos-mesh/api/v1alpha1
BUILD_INDOCKER_ARG := --env IN_DOCKER=1 --volume $(ROOT):/mnt --user $(shell id -u):$(shell id -g)
$(info $g11)
define MA
aa:
	$(info xxx)
endef

$(eval $(call MA) )
l1:
  $(info $PACKAGE_LIST)
  $(info $BUILD_INDOCKER_ARG)
