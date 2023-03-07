# Stage 1: install dependencies
FROM node:lts-alpine AS deps
WORKDIR /app
COPY ./frontend/package*.json ./
ARG NODE_ENV
ENV NODE_ENV $NODE_ENV
RUN npm install

# Stage 2: build
FROM node:lts-alpine AS builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY ./frontend/src ./src
COPY ./frontend/public ./public
COPY ./frontend/package.json ./package.json
COPY ./frontend/next.config.mjs ./next.config.mjs
COPY ./frontend/tsconfig.json ./tsconfig.json
COPY ./frontend/.env ./.env
COPY ./frontend/*.cjs ./
RUN npm run build

# Stage 3: run
FROM node:lts-alpine AS runner
WORKDIR /app
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./
CMD ["npm", "run", "start"]
