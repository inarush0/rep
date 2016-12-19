FROM golang:alpine

# Create the directory where the application will reside
RUN mkdir /app

ADD rep /app/rep
ADD templates /app/templates

WORKDIR /app

EXPOSE 8080

ENTRYPOINT /app/rep
