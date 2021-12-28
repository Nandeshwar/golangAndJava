FROM openjdk:15-jdk-alpine AS builder
RUN apk add --no-cache maven

WORKDIR /quote
COPY . .

RUN mvn clean
RUN /quote/mvnw package

FROM openjdk:15-jdk-alpine
ARG JAR_FILE=/quote/target/quote-0.1.0.jar
COPY --from=builder ${JAR_FILE} quote.jar

ENTRYPOINT ["java", "-jar", "/quote.jar"]