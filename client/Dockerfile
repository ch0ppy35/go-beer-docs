FROM node:18 as build

WORKDIR /app
COPY package*.json /app/
RUN npm install
COPY ./ /app/
RUN npm run build

FROM nginx:stable-alpine as dist
COPY --from=build /app/build/ /usr/share/nginx/html