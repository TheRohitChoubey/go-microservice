FROM node:14.17.1 as node
WORKDIR /app
COPY . .
RUN npm install
RUN npm run build --prod

FROM nginx:alpine
COPY --from=node /app/dist/vmWare-assignment /usr/share/nginx/html