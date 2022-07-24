FROM node:latest

RUN apt -y update
RUN apt -y upgrade
RUN npm install -g npm@latest
ADD ./my-app /my-app
WORKDIR /my-app
ENV NODE_OPTIONS=--openssl-legacy-provider
RUN npm i
RUN mkdir /build
RUN npm run-script build
RUN mv /my-app/build/* /build
