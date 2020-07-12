.PHONY: generate
generate: generate-greet generate-calculator

.PHONY: generate-greet
generate-greet:
	@protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

.PHONY: generate-calculator
generate-calculator:
	@protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
