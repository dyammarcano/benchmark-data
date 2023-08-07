FROM openjdk:17-buster AS build

# Set the working directory inside the container
WORKDIR /app

COPY . .

# Build the application

RUN ./mvnw clean package -DskipTests

# Path: Dockerfile
FROM openjdk:17-buster

# Copy the JAR file from the target directory to the container
COPY --from=build /app/target/*.jar app.jar

# Expose the port that the application will listen on
EXPOSE 8080

# Command to run the application
CMD ["java", "-jar", "app.jar"]
