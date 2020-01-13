FROM node:10.15.1-alpine as builder
WORKDIR /app/client
COPY . /app

RUN apk update && \
    apk add git

RUN yarn install --production
RUN yarn build

FROM node:10.15.1-alpine
WORKDIR /app/client
COPY --from=builder /app /app

CMD ["yarn", "start"]
