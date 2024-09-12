# Stage 1: Build the React application
FROM node:20 AS build

# Set working directory
WORKDIR /app

# Install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy source files
COPY . .

# Build the React application
RUN npm run build

# Stage 2: Serve the React application with Nginx
FROM nginx:alpine

# Copy the built React application from the previous stage
COPY --from=build /app/build /usr/share/nginx/html

# Expose port 80 for the application
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]
