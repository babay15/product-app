#Getting latest golang from docker hub
FROM golang:latest

#Set maintainer
LABEL maintainer="babaymxc"

#Set working directory inside container
WORKDIR /app

#Copy all dependency
COPY go.mod .

#Copy Module
COPY go.sum .

#Download dependency
RUN go mod download

#Copy all files in the project
COPY . .

#Set environment
ENV PORT 5000
ENV JWT_SECRET "babay15"

#Build application
RUN go build

#After succesfully build the app, remove all source files
RUN find . -name "*.go" -type f -delete

#Expose port
EXPOSE $PORT

#Run application
CMD ["./product-app"]