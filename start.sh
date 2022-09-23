docker build -t docker-my-server-ms -f Dockerfile .
docker run -d -p 8080:3001 --name my-server docker-my-server-ms
echo
echo "Running my-server on port 8080..."
echo