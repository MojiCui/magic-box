#!/bin/bash

source /opt/$PROJECT_NAME/docker/preset.sh
source /opt/$PROJECT_NAME/docker/loc.env
source /opt/$PROJECT_NAME/docker/postset.sh

export PROJECT_NAME=${PROJECT_NAME}
export APP_NAME=${APP_NAME}

# 启动服务
exec /opt/$PROJECT_NAME/$APP_NAME serve
