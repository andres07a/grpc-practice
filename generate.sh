#!/bin/bash

# setting execute permission
#$ chmod a+x generate.sh

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.