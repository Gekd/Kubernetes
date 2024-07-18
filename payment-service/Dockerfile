FROM node:14-alpine

WORKDIR /app

COPY package.json ./
COPY tsconfig.json ./
RUN npm install

COPY src/ ./src/

EXPOSE 3000

RUN npm run build

CMD ["node", "dist/index.js"]
