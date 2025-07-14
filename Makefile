# Makefile for challenge-clean-architecture

# Default target when you just type 'make'
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make up         - Start all services using Docker Compose"
	@echo "  make down      - Stop and remove all Docker containers"
	@echo "  make run       - Run the application directly with Go"

# Start all services using Docker Compose
.PHONY: run
run:
	docker-compose up --build
