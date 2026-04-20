
ARG BASE_IMAGE=hub-mirror.wps.cn/wpsplusopen/golang-alpine:v1.19
ARG BASE_BUILD_IMAGE=hub-mirror.wps.cn/wpsplusopen/golang-build:v1.19

FROM ${BASE_BUILD_IMAGE} as golang-build

ARG GIT_BRANCH

# 根据CI配置克隆目录修改PROJECT_NAME参数
ENV PROJECT_NAME=magic-box
ENV APP_NAME=magic-box
ENV GIT_BRANCH=${GIT_BRANCH}

WORKDIR /opt/${PROJECT_NAME}/${APP_NAME}
ADD ./${PROJECT_NAME} ./
RUN make gen branch=${GIT_BRANCH} && make build

FROM ${BASE_IMAGE}

# 根据CI配置克隆目录修改PROJECT_NAME参数
ENV PROJECT_NAME=magic-box
ENV APP_NAME=magic-box

WORKDIR /opt/${PROJECT_NAME}
COPY --from=golang-build  /opt/${PROJECT_NAME}/${APP_NAME}/docker ./docker
COPY --from=golang-build  /opt/${PROJECT_NAME}/${APP_NAME}/${APP_NAME} ./

RUN chmod +x ./docker/start.sh
RUN chmod +x ./docker/livenessProbe.sh
RUN chmod +x ./docker/readinessProbe.sh

RUN cp /opt/${PROJECT_NAME}/docker/livenessProbe.sh /usr/local/bin/livenessProbe.sh
RUN cp /opt/${PROJECT_NAME}/docker/readinessProbe.sh /usr/local/bin/readinessProbe.sh


EXPOSE 80 8081
ENTRYPOINT ["/bin/bash", "-c","/opt/${PROJECT_NAME}/docker/start.sh"]
