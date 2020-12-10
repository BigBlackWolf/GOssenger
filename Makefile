check:
ifeq ($(shell which go), )
	@echo "'go' is not installed"
else
	@echo "'go' is installed"
endif
ifeq ($(shell which mongo), )
	@echo "'mongo' is not installed"
else
	@echo "'mongo' is installed"
endif
ifeq ($(shell which np2m), )
	@echo "'npm' is not installed"
else
	@echo "'npm' is installed"
endif