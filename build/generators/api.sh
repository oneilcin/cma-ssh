#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

THIS_DIRECTORY=$(dirname "${BASH_SOURCE}")
PROJECT_DIRECTORY=${THIS_DIRECTORY}/../..
SWAGGER_DESTINATION=assets/generated/swagger

echo "Ensuring Swagger Asset Direcotry ( ${SWAGGER_DESTINATION} ) Exists..."
mkdir -p ${PROJECT_DIRECTORY}/${SWAGGER_DESTINATION}

export PATH=${PROJECT_DIRECTORY}/bin:$PATH
"${PROJECT_DIRECTORY}/bin/protoc/bin/protoc" "${PROJECT_DIRECTORY}/api/api.proto" \
  -I "${PROJECT_DIRECTORY}/api" \
  -I "${PROJECT_DIRECTORY}/api/third_party/" \
  --go_out=plugins=grpc:"${PROJECT_DIRECTORY}/" \
  --grpc-gateway_out=logtostderr=true:"${PROJECT_DIRECTORY}" \
  --swagger_out=logtostderr=true:"${PROJECT_DIRECTORY}/${SWAGGER_DESTINATION}" \
  --doc_out "${PROJECT_DIRECTORY}/docs/api-generated" \
  --doc_opt=markdown,api.md
