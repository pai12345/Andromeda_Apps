# Python base image
FROM python:alpine3.17

# Create user and group
RUN addgroup app && adduser -S -G app app

# Create working directory
WORKDIR /app

# Copy contents to app
COPY . .

# Install dependencies
RUN pip install -r requirements.txt

# Set user
USER app

# Expose Port
EXPOSE 8080

# Execute
ENTRYPOINT [ "python", "index.py" ]
